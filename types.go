package telegram

// User represents a Telegram user or bot.
//
// https://core.telegram.org/bots/api#user.
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`                   // Optional.
	Username                string `json:"username,omitempty"`                    // Optional.
	LanguageCode            string `json:"language_code,omitempty"`               // Optional.
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`             // Optional.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"` // Optional.
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`     // Optional.
}

// Chat represents a chat.
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title,omitempty"`               // Optional.
	Username         string           `json:"username,omitempty"`            // Optional.
	FirstName        string           `json:"first_name,omitempty"`          // Optional.
	LastName         string           `json:"last_name,omitempty"`           // Optional.
	Photo            *ChatPhoto       `json:"photo,omitempty"`               // Optional.
	Bio              string           `json:"bio,omitempty"`                 // Optional.
	Description      string           `json:"description,omitempty"`         // Optional.
	InviteLink       string           `json:"invite_link,omitempty"`         // Optional.
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`      // Optional.
	Permissions      *ChatPermissions `json:"permissions,omitempty"`         // Optional.
	SlowModeDelay    int              `json:"slow_mode_delay,omitempty"`     // Optional.
	StickerSetName   string           `json:"sticker_set_name,omitempty"`    // Optional.
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"` // Optional.
	LinkedChatID     int              `json:"linked_chat_id,omitempty"`      // Optional.
	Location         *ChatLocation    `json:"location,omitempty"`            // Optional.
}

// Message represents a message.
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID               int                      `json:"message_id"`
	From                    *User                    `json:"from,omitempty"`        // Optional.
	SenderChat              *Chat                    `json:"sender_chat,omitempty"` // Optional.
	Date                    int                      `json:"date"`
	Chat                    *Chat                    `json:"chat"`
	ForwardFrom             *User                    `json:"forward_from,omitempty"`              // Optional.
	ForwardFromChat         *Chat                    `json:"forward_from_chat,omitempty"`         // Optional.
	ForwardFromMessageID    int                      `json:"forward_from_message_id,omitempty"`   // Optional.
	ForwardSignature        string                   `json:"forward_signature,omitempty"`         // Optional.
	ForwardSenderName       string                   `json:"forward_sender_name,omitempty"`       // Optional.
	ForwardDate             int                      `json:"forward_date,omitempty"`              // Optional.
	ReplyToMessage          *Message                 `json:"reply_to_message,omitempty"`          // Optional.
	ViaBot                  *User                    `json:"via_bot,omitempty"`                   // Optional.
	EditDate                int                      `json:"edit_date,omitempty"`                 // Optional.
	MediaGroupID            string                   `json:"media_group_id,omitempty"`            // Optional.
	AuthorSignature         string                   `json:"author_signature,omitempty"`          // Optional.
	Text                    string                   `json:"text,omitempty"`                      // Optional.
	Entities                []*MessageEntity         `json:"entities,omitempty"`                  // Optional.
	Animation               *Animation               `json:"animation,omitempty"`                 // Optional.
	Audio                   *Audio                   `json:"audio,omitempty"`                     // Optional.
	Document                *Document                `json:"document,omitempty"`                  // Optional.
	Photo                   []*PhotoSize             `json:"photo,omitempty"`                     // Optional.
	Sticker                 *Sticker                 `json:"sticker,omitempty"`                   // Optional.
	Video                   *Video                   `json:"video,omitempty"`                     // Optional.
	VideoNote               *VideoNote               `json:"video_note,omitempty"`                // Optional.
	Voice                   *Voice                   `json:"voice,omitempty"`                     // Optional.
	Caption                 string                   `json:"caption,omitempty"`                   // Optional.
	CaptionEntities         []*MessageEntity         `json:"caption_entities,omitempty"`          // Optional.
	Contact                 *Contact                 `json:"contact,omitempty"`                   // Optional.
	Dice                    *Dice                    `json:"dice,omitempty"`                      // Optional.
	Game                    *Game                    `json:"game,omitempty"`                      // Optional.
	Poll                    *Poll                    `json:"poll,omitempty"`                      // Optional.
	Venue                   *Venue                   `json:"venue,omitempty"`                     // Optional.
	Location                *Location                `json:"location,omitempty"`                  // Optional.
	NewChatMembers          []*User                  `json:"new_chat_members,omitempty"`          // Optional.
	LeftChatMember          *User                    `json:"left_chat_member,omitempty"`          // Optional.
	NewChatTitle            string                   `json:"new_chat_title,omitempty"`            // Optional.
	NewChatPhoto            []*PhotoSize             `json:"new_chat_photo,omitempty"`            // Optional.
	DeleteChatPhoto         bool                     `json:"delete_chat_photo,omitempty"`         // Optional.
	GroupChatCreated        bool                     `json:"group_chat_created,omitempty"`        // Optional.
	SupergroupChatCreated   bool                     `json:"supergroup_chat_created,omitempty"`   // Optional.
	ChannelChatCreated      bool                     `json:"channel_chat_created,omitempty"`      // Optional.
	MigrateToChatID         int                      `json:"migrate_to_chat_id,omitempty"`        // Optional.
	MigrateFromChatID       int                      `json:"migrate_from_chat_id,omitempty"`      // Optional.
	PinnedMessage           *Message                 `json:"pinned_message,omitempty"`            // Optional.
	Invoice                 *Invoice                 `json:"invoice,omitempty"`                   // Optional.
	SuccessfulPayment       *SuccessfulPayment       `json:"successful_payment,omitempty"`        // Optional.
	ConnectedWebsite        string                   `json:"connected_website,omitempty"`         // Optional.
	PassportData            *PassportData            `json:"passport_data,omitempty"`             // Optional.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"` // Optional.
	ReplyMarkup             *InlineKeyboardMarkup    `json:"reply_markup,omitempty"`              // Optional.
}

// MessageID represents a unique message identifier.
//
// https://core.telegram.org/bots/api#messageid
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url,omitempty"`      // Optional.
	User     *User  `json:"user,omitempty"`     // Optional.
	Language string `json:"language,omitempty"` // Optional.
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"` // Optional.
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
//
// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`     // Optional.
	FileName     string     `json:"file_name,omitempty"` // Optional.
	MIMEType     string     `json:"mime_type,omitempty"` // Optional.
	FileSize     int        `json:"file_size,omitempty"` // Optional.
}

// Audio represents an audio file to be treated as music by the Telegram clients.
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"` // Optional.
	Title        string     `json:"title,omitempty"`     // Optional.
	FileName     string     `json:"file_name,omitempty"` // Optional.
	MIMEType     string     `json:"mime_type,omitempty"` // Optional.
	FileSize     int        `json:"file_size,omitempty"` // Optional.
	Thumb        *PhotoSize `json:"thumb,omitempty"`     // Optional.
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`     // Optional.
	FileName     string     `json:"file_name,omitempty"` // Optional.
	MIMEType     string     `json:"mime_type,omitempty"` // Optional.
	FileSize     int        `json:"file_size,omitempty"` // Optional.
}

// Video represents a video file.
//
// https://core.telegram.org/bots/api#video
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`     // Optional.
	FileName     string     `json:"file_name,omitempty"` // Optional.
	MIMEType     string     `json:"mime_type,omitempty"` // Optional.
	FileSize     int        `json:"file_size,omitempty"` // Optional.
}

// VideoNote represents a video message.
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`     // Optional.
	FileSize     int        `json:"file_size,omitempty"` // Optional.
}

// Voice represents a voice note.
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     string `json:"duration"`
	MIMEType     string `json:"mime_type,omitempty"` // Optional.
	FileSize     string `json:"file_size,omitempty"` // Optional.
}

// Contact represents a phone contact.
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"` // Optional.
	UserID      int    `json:"user_id,omitempty"`   // Optional.
	VCard       string `json:"vcard,omitempty"`     // Optional.
}

// Dice represents an animated emoji that displays a random value.
//
// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption contains information about one answer option in a poll.
//
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
//
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll contains information about a poll.
//
// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`    // Optional.
	Explanation           string           `json:"explanation,omitempty"`          // Optional.
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"` // Optional.
	OpenPeriod            int              `json:"open_period,omitempty"`          // Optional.
	CloseDate             int              `json:"close_date,omitempty"`           // Optional.
}

// Location represents a point on the map.
//
// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional.
	LivePeriod           int     `json:"live_period,omitempty"`            // Optional.
	Heading              int     `json:"heading,omitempty"`                // Optional.
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"` // Optional.
}

// Venue represents a venue.
//
// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`     // Optional.
	FoursquareType  string    `json:"foursquare_type,omitempty"`   // Optional.
	GooglePlaceID   string    `json:"google_place_id,omitempty"`   // Optional.
	GooglePlaceType string    `json:"google_place_type,omitempty"` // Optional.
}

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
//
// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

// UserProfilePhotos represent a user's profile pictures.
//
// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling GetFile.
//  Maximum file size to download is 20 MB
//
// https://core.telegram.org/bots/api#file
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     string `json:"file_size,omitempty"` // Optional.
	FilePath     string `json:"file_path,omitempty"` // Optional.
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`   // Optional.
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"` // Optional.
	Selective       bool                `json:"selective,omitempty"`         // Optional.
}

// KeyboardButton represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields request_contact, request_location, and request_poll are mutually exclusive.
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`  // Optional.
	RequestLocation bool                    `json:"request_location,omitempty"` // Optional.
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`     // Optional.
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
//
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"` // Optional.
}

// ReplyKeyboardRemove, upon receiving a message with this object,
// Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
//
// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"` // Optional.
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
//
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`                              // Optional.
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`                        // Optional.
	CallbackData                 string        `json:"callback_data,omitempty"`                    // Optional.
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`              // Optional.
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"` // Optional.
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`                    // Optional.
	Pay                          bool          `json:"pay,omitempty"`                              // Optional.
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in.
//
// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`         // Optional.
	BotUsername        string `json:"bot_username,omitempty"`         // Optional.
	RequestWriteAccess bool   `json:"request_write_access,omitempty"` // Optional.
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`           // Optional.
	InlineMessageID string   `json:"inline_message_id,omitempty"` // Optional.
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`            // Optional.
	GameShortName   string   `json:"game_short_name,omitempty"` // Optional.
}

// ForceReply, upon receiving a message with this object,
// Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
//
// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"` // Optional.
}

// ChatPhoto represents a chat photo.
//
// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatMember contains information about one member of a chat.
//
// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`              // Optional.
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`              // Optional.
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // Optional.
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // Optional.
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // Optional.
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // Optional.
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // Optional.
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // Optional.
	CahChangeInfo         bool   `json:"can_change_info,omitempty"`           // Optional.
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // Optional.
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // Optional.
	IsMember              bool   `json:"is_member,omitempty"`                 // Optional.
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // Optional.
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`   // Optional.
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`            // Optional.
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // Optional.
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // Optional.
	UntilDate             int    `json:"until_date,omitempty"`                // Optional.
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
//
// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`         // Optional.
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`   // Optional.
	CanSendPolls          bool `json:"can_send_polls,omitempty"`            // Optional.
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`   // Optional.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"` // Optional.
	CanChangeInfo         bool `json:"can_change_info,omitempty"`           // Optional.
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`          // Optional.
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`          // Optional.
}

// ChatLocation represents a location to which a chat is connected.
//
// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// BotCommand represents a bot command.
//
// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ResponseParameters contains information about why a request was unsuccessful.
//
// https://core.telegram.org/bots/api#responseparameters
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"` // Optional.
	RetryAfter      int `json:"retry_after,omitempty"`        // Optional.
}

// InputMedia represents the content of a media message to be sent. It should be one of
//
//  InputMediaAnimation
//  InputMediaDocument
//  InputMediaAudio
//  InputMediaPhoto
//  InputMediaVideo
//
// https://core.telegram.org/bots/api#inputmedia
// TODO: should InputMedia be a type?

// InputMediaPhoto represents a photo to be sent.
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Caption         string           `json:"caption,omitempty"`          // Optional.
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional.
}

// InputMediaVideo represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// Thumb *InputFile or string `json:"thumb,omitempty"` // TODO: handle the inconsistent type
	Caption           string           `json:"caption,omitempty"`            // Optional.
	ParseMode         string           `json:"parse_mode,omitempty"`         // Optional.
	CaptionEntities   []*MessageEntity `json:"caption_entities,omitempty"`   // Optional.
	Width             int              `json:"width,omitempty"`              // Optional.
	Height            int              `json:"height,omitempty"`             // Optional.
	Duration          int              `json:"duration,omitempty"`           // Optional.
	SupportsStreaming bool             `json:"supports_streaming,omitempty"` // Optional.
}

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// Thumb *InputFile or string `json:"thumb,omitempty"` // TODO: handle the inconsistent type
	Caption         string           `json:"caption,omitempty"`          // Optional.
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional.
	Width           int              `json:"width,omitempty"`            // Optional.
	Height          int              `json:"height,omitempty"`           // Optional.
	Duration        int              `json:"duration,omitempty"`         // Optional.
}

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// Thumb *InputFile or string `json:"thumb,omitempty"` // TODO: handle the inconsistent type
	Caption         string           `json:"caption,omitempty"`          // Optional.
	ParseMode       string           `json:"parse_mode,omitempty"`       // Optional.
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"` // Optional.
	Duration        int              `json:"duration,omitempty"`         // Optional.
	Performer       string           `json:"performer,omitempty"`        // Optional.
	Title           string           `json:"title,omitempty"`            // Optional.
}

// InputMediaDocument represents a general file to be sent.
//
// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// Thumb *InputFile or string `json:"thumb,omitempty"` // TODO: handle the inconsistent type
	Caption                     string           `json:"caption,omitempty"`                        // Optional.
	ParseMode                   string           `json:"parse_mode,omitempty"`                     // Optional.
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`               // Optional.
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"` // Optional.
}

// InputFile represents the contents of a file to be uploaded.
// Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
//
// https://core.telegram.org/bots/api#inputfile
// TODO: should InputFile be a type?
