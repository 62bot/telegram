package telegram

// InlineQuery represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"` // Optional.
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// InlineQueryResult represents one result of an inline query.
// Telegram clients currently support results of the following 20 types:
//
// InlineQueryResultCachedAudio
// InlineQueryResultCachedDocument
// InlineQueryResultCachedGIF
// InlineQueryResultCachedMPEG4GIF
// InlineQueryResultCachedPhoto
// InlineQueryResultCachedSticker
// InlineQueryResultCachedVideo
// InlineQueryResultCachedVoice
// InlineQueryResultArticle
// InlineQueryResultAudio
// InlineQueryResultContact
// InlineQueryResultGame
// InlineQueryResultDocument
// InlineQueryResultGIF
// InlineQueryResultLocation
// InlineQueryResultMPEG4GIF
// InlineQueryResultPhoto
// InlineQueryResultVenue
// InlineQueryResultVideo
// InlineQueryResultVoice
//
// https://core.telegram.org/bots/api#inlinequeryresult
// TODO: should InputMedia be a type?

// InlineQueryResultArticle represents a link to an article or web page.
//
// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent interface{}           `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional.
	URL                 string                `json:"url,omitempty"`          // Optional.
	HideURL             bool                  `json:"hide_url,omitempty"`     // Optional.
	Description         string                `json:"description,omitempty"`  // Optional.
	ThumbURL            string                `json:"thumb_url,omitempty"`    // Optional.
	ThumbWidth          int                   `json:"thumb_width,omitempty"`  // Optional.
	ThumbHeight         int                   `json:"thumb_height,omitempty"` // Optional.
}

// InlineQueryResultPhoto represents a link to a photo.
// By default,this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//
// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`           // Optional.
	PhotoHeight         int                   `json:"photo_height,omitempty"`          // Optional.
	Title               string                `json:"title,omitempty"`                 // Optional.
	Description         string                `json:"description,omitempty"`           // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultGIF represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGIF struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GIFURL              string                `json:"gif_url"`
	GIFWidth            int                   `json:"gif_width,omitempty"`    // Optional.
	GIFHeight           int                   `json:"gif_height,omitempty"`   // Optional.
	GIFDuration         int                   `json:"gif_duration,omitempty"` // Optional.
	ThumbURL            string                `json:"thumb_url"`
	ThumbMIMEType       string                `json:"thumb_mime_type,omitempty"`       // Optional.
	Title               string                `json:"title,omitempty"`                 // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultMPEG4GIF represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMPEG4GIF struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	MPEG4URL            string                `json:"mpeg4_url"`
	MPEG4Width          int                   `json:"mpeg4_width,omitempty"`    // Optional.
	MPEG4Height         int                   `json:"mpeg4_height,omitempty"`   // Optional.
	MPEG4Duration       int                   `json:"mpeg4_duration,omitempty"` // Optional.
	ThumbURL            string                `json:"thumb_url"`
	ThumbMIMEType       string                `json:"thumb_mime_type,omitempty"`       // Optional.
	Title               string                `json:"title,omitempty"`                 // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MIMEType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	VideoWidth          int                   `json:"video_width,omitempty"`           // Optional.
	VideoHeight         int                   `json:"video_height,omitempty"`          // Optional.
	VideoDuration       int                   `json:"video_duration,omitempty"`        // Optional.
	Description         string                `json:"description,omitempty"`           // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultAudio represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//
// https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	Performer           string                `json:"performer,omitempty"`             // Optional.
	AudioDuration       int                   `json:"audio_duration,omitempty"`        // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
//
// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	VoiceDuration       int                   `json:"voice_duration,omitempty"`        // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultDocument represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
//
// https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`          // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`       // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"` // Optional.
	DocumentURL         string                `json:"document_url"`
	MIMEType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`           // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional.
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional.
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional.
}

// InlineQueryResultLocation represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
//
// https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`    // Optional.
	LivePeriod           int                   `json:"live_period,omitempty"`            // Optional.
	Heading              int                   `json:"heading,omitempty"`                // Optional.
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"` // Optional.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`           // Optional.
	InputMessageContent  interface{}           `json:"input_message_content,omitempty"`  // Optional.
	ThumbURL             string                `json:"thumb_url,omitempty"`              // Optional.
	ThumbWidth           int                   `json:"thumb_width,omitempty"`            // Optional.
	ThumbHeight          int                   `json:"thumb_height,omitempty"`           // Optional.
}

// InlineQueryResultVenue represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
//
// https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`         // Optional.
	FoursquareType      string                `json:"foursquare_type,omitempty"`       // Optional.
	GooglePlaceID       string                `json:"google_place_id,omitempty"`       // Optional.
	GooglePlaceType     string                `json:"google_place_type,omitempty"`     // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional.
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional.
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional.
}

// InlineQueryResultContact represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
//
// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`             // Optional.
	VCard               string                `json:"vcard,omitempty"`                 // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
	ThumbURL            string                `json:"thumb_url,omitempty"`             // Optional.
	ThumbWidth          int                   `json:"thumb_width,omitempty"`           // Optional.
	ThumbHeight         int                   `json:"thumb_height,omitempty"`          // Optional.
}

// InlineQueryResultGame represents a Game.
//
// https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"` // Optional.
}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               string                `json:"title,omitempty"`                 // Optional.
	Description         string                `json:"description,omitempty"`           // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedGIF represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGIF struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GIFFileID           string                `json:"gif_file_id"`
	Title               string                `json:"title,omitempty"`                 // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedMPEG4GIF represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMPEG4GIF struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	MPEG4FileID         string                `json:"mpeg4_file_id"`
	Title               string                `json:"title,omitempty"`                 // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
//
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`           // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`           // Optional.
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`               // Optional.
	ParseMode           string                `json:"parse_mode,omitempty"`            // Optional.
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`      // Optional.
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Optional.
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // Optional.
}

// InputMessageContent represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 4 types:
//
// InputTextMessageContent
// InputLocationMessageContent
// InputVenueMessageContent
// InputContactMessageContent
//
// https://core.telegram.org/bots/api#inputmessagecontent
// TODO: should InputMessageContent be a type?

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
//
// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional.
	Entities              []*MessageEntity `json:"entities,omitempty"`                 // Optional.
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"` // Optional.
}

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
//
// https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional.
	LivePeriod           int     `json:"live_period,omitempty"`            // Optional.
	Heading              int     `json:"heading,omitempty"`                // Optional.
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"` // Optional.
}

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
//
// https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`     // Optional.
	FoursquareType  string  `json:"foursquare_type,omitempty"`   // Optional.
	GooglePlaceID   string  `json:"google_place_id,omitempty"`   // Optional.
	GooglePlaceType string  `json:"google_place_type,omitempty"` // Optional.
}

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
//
// https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber float64 `json:"phone_number"`
	FirstName   float64 `json:"first_name"`
	LastName    string  `json:"last_name,omitempty"` // Optional.
	VCard       string  `json:"vcard,omitempty"`     // Optional.
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`          // Optional.
	InlineMessageID string    `json:"inline_message_id,omitempty"` // Optional.
	Query           string    `json:"query"`
}
