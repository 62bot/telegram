package telegram

// Sticker represents a sticker.
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`         // Optional.
	Emoji        string        `json:"emoji,omitempty"`         // Optional.
	SetName      string        `json:"set_name,omitempty"`      // Optional.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"` // Optional.
	FileSize     int           `json:"file_size,omitempty"`     // Optional.
}

// StickerSet represents a sticker set.
//
// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"` // Optional.
}

// MaskPosition describes the position on faces where a mask should be placed by default.
//
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}
