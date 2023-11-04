package responses

type GiftsReply struct {
	Error    *Error `json:"error"`
	Response Gifts  `json:"response,omitempty"`
}

type Gifts struct {
	Error *Error `json:"error"`
	Count int    `json:"count,omitempty"`
	Gifts []Gift `json:"items,omitempty"`
}

type Gift struct {
	ID      int      `json:"id,omitempty"`
	UserID  int      `json:"from_id,omitempty"`
	Message string   `json:"message,omitempty"`
	Date    int      `json:"date,omitempty"`
	Gift    GiftInfo `json:"gift,omitempty"`
	Privacy int      `json:"privacy,omitempty"`
}

type GiftInfo struct {
	ID       int    `json:"id,omitempty"`
	Photo48  string `json:"thumb_48,omitempty"`
	Photo96  string `json:"thumb_96,omitempty"`
	Photo256 string `json:"thumb_256,omitempty"`
}
