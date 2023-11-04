package update

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions,omitempty"`
	Keyboard       bool     `json:"keyboard,omitempty"`
	InlineKeyboard bool     `json:"inline_keyboard,omitempty"`
	Carousel       bool     `json:"carousel,omitempty"`
	LanguageID     int      `json:"lang_id,omitempty"`
}
