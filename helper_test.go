package telegram_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/62bot/telegram"
)

func TestMain(m *testing.M) {
	// setup
	baseDir := "testdata"
	decodeTestCases := func(methodName, filePath string) testCaseStorage {
		f, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		var v map[string]testCase
		if err := json.NewDecoder(f).Decode(&v); err != nil {
			panic(err)
		}

		for caseName, tc := range v {
			v[caseName] = testCase{
				methodName: methodName,
				caseName:   caseName,
				StatusCode: tc.StatusCode,
				Params:     tc.Params,
				Body:       tc.Body,
			}
			tc.caseName = caseName
		}

		return testCaseStorage{
			methodName: methodName,
			data:       v,
		}
	}

	fis, err := ioutil.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		var (
			fileName   = fi.Name()
			ext        = path.Ext(fileName)
			methodName = strings.TrimSuffix(fileName, ext)
		)

		if ext != ".json" {
			continue
		}

		testFixture[methodName] = decodeTestCases(methodName, path.Join(baseDir, fileName))
	}

	// pre-assign common testCase
	testCases := testFixture.get("getMe")
	unauthorizedCase = testCases.get("unauthorized")
	authorizedCase = testCases.get("authorized")

	// run the test
	code := m.Run()

	// cleanup
	os.Exit(code)
}

var (
	testFixture = make(testFixtureStorage)

	unauthorizedCase testCase
	authorizedCase   testCase

	validTestToken   = "123456:valid-token"
	invalidTestToken = "123456:invalid-token"
)

// testFixtureStorage format: map[{METHOD_NAME}]testCases.
type testFixtureStorage map[string]testCaseStorage

func (storage testFixtureStorage) get(methodName string) testCaseStorage {
	testCases, ok := storage[methodName]
	if !ok {
		panic(fmt.Errorf("no test fixture for method %q", methodName))
	}
	return testCases
}

// testCaseStorage format: map[{CASE_NAME}]testCase.
type testCaseStorage struct {
	methodName string
	data       map[string]testCase
}

func (storage *testCaseStorage) get(caseName string) testCase {
	tc, ok := storage.data[caseName]
	if !ok {
		panic(fmt.Errorf("no test data for case %q in method %q", caseName, storage.methodName))
	}
	return tc
}

type testCase struct {
	methodName string
	caseName   string

	StatusCode int
	Params     url.Values
	Body       json.RawMessage
}

func (tc *testCase) UnmarshalJSON(b []byte) error {
	type TC struct {
		StatusCode int             `json:"status_code"`
		Params     string          `json:"params"`
		Body       json.RawMessage `json:"body"`
	}

	var temp TC
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	params, err := url.ParseQuery(temp.Params)
	if err != nil {
		return err
	}

	tc.StatusCode = temp.StatusCode
	tc.Params = params
	tc.Body = temp.Body

	return nil
}

func (tc *testCase) decodeResponse(v interface{}) {
	var resp telegram.Response
	if err := json.Unmarshal(tc.Body, &resp); err != nil {
		panic(fmt.Errorf("%s/%s: %v", tc.methodName, tc.caseName, err))
	}
	if err := json.Unmarshal(resp.Result, v); err != nil {
		panic(fmt.Errorf("%s/%s: %v", tc.methodName, tc.caseName, err))
	}
}

type roundTripperTestFunc func(methodName string, params url.Values) *http.Response

func (f roundTripperTestFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	var (
		urlPath    = r.URL.Path
		methodName = path.Base(urlPath)
		params     = r.URL.Query()
		token      = strings.TrimPrefix(strings.TrimSuffix(urlPath, "/"+methodName), "/bot")
	)

	if token != validTestToken {
		return newHTTPResponse(unauthorizedCase.StatusCode, unauthorizedCase.Body), nil
	}

	if token == validTestToken && methodName == "getMe" {
		return newHTTPResponse(authorizedCase.StatusCode, authorizedCase.Body), nil
	}

	return f(methodName, params), nil
}

// newTestClient returns an http client for test.
func newTestClient(fn roundTripperTestFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// newResponse returns an http response with a specified status code
// and JSON encoded body from v.
func newHTTPResponse(statusCode int, v []byte) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewReader(v)),
	}
}
