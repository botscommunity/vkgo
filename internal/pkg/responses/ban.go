package responses

type Ban struct {
	Error *Error `json:"error"`
	Code  int    `json:"response"`
}

type BannedReply struct {
	Error    *Error `json:"error"`
	Response Banned `json:"response"`
}

type Banned struct {
	Error *Error `json:"error"`
	Count int    `json:"count,omitempty"`
	Bans  []int  `json:"items,omitempty"`
	Users []User `json:"profiles,omitempty"`
}
