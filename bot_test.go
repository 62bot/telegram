package telegram_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/62bot/telegram"
	"github.com/billyzaelani/is"
)

// TestNewBot tests creation of new bot. In addition it's also test method getMe.
func TestNewBot(t *testing.T) {
	is := is.New(t)

	client := newTestClient(func(methodName string, params url.Values) *http.Response {
		return newHTTPResponse(authorizedCase.StatusCode, authorizedCase.Body)
	})

	_, err := telegram.NewBot(invalidTestToken, telegram.SetClient(client))
	var botError *telegram.BotError
	is.ErrorAs(err, &botError)
	is.Equal(botError.Code, http.StatusUnauthorized)

	bot, err := telegram.NewBot(validTestToken, telegram.SetClient(client))
	is.NoError(err)

	var wantUser telegram.User
	authorizedCase.decodeResponse(&wantUser)

	is.Equal(bot.User.ID, wantUser.ID)
	is.Equal(bot.User.IsBot, wantUser.IsBot)
	is.Equal(bot.User.FirstName, wantUser.FirstName)
	is.Equal(bot.User.Username, wantUser.Username)
	is.Equal(bot.User.CanJoinGroups, wantUser.CanJoinGroups)
	is.Equal(bot.User.CanReadAllGroupMessages, wantUser.CanReadAllGroupMessages)
	is.Equal(bot.User.SupportsInlineQueries, wantUser.SupportsInlineQueries)
}

// TestMethod tests all method listed in method_test.go file
func TestMethod(t *testing.T) {
	tests := map[string]methodTestFunc{
		// test getMe in TestNewBot
		// don't test getMe here

		// updates_test
		"getUpdates/ok":                    getUpdatesOK,
		"getUpdates/with_params":           getUpdatesWithParams,
		"getUpdates/conflict_with_webhook": getUpdatesConflictWithWebhook,
		"setWebhook/ok":                    setWebhookOK,
		"setWebhook/with_params":           setWebhookWithParams,
		"setWebhook/error_not_https":       setWebhookErrorNotHTTPS,
		"deleteWebhook/ok":                 deleteWebhookOK,
		"getWebhookInfo/ok":                getWebhookInfoOK,

		// methods_test
		"sendMessage/ok": sendMessageOK,
	}

	for name, f := range tests {
		name := name
		f := f
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := strings.Split(name, "/")
			if len(s) != 2 {
				t.Fatalf("test name %q should in form {methodName}/{caseName}", s)
			}

			var (
				is             = is.New(t)
				wantMethodName = s[0]
				wantCaseName   = s[1]
			)

			testCases := testFixture.get(wantMethodName)
			tc := testCases.get(wantCaseName)

			client := newTestClient(func(methodName string, params url.Values) *http.Response {
				is.Equal(methodName, wantMethodName) // bot calling unwanted method
				is.Equal(params, tc.Params)          // bot calling method with wrong params
				return newHTTPResponse(tc.StatusCode, tc.Body)
			})

			bot, err := telegram.NewBot(validTestToken, telegram.SetClient(client))
			is.NoError(err) // new bot should not error in this test

			f(is, bot)
		})
	}
}
