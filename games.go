package telegram

// Game represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
//
// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text,omitempty"`          // Optional.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"` // Optional.
	Animation    *Animation       `json:"animation,omitempty"`     // Optional.
}

// CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// GameHighScore represents one row of the high scores table for a game.
//
// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
