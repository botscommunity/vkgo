package responses

type LongPollServerReply struct {
	Error    *Error         `json:"error"`
	Response LongPollServer `json:"response"`
}

type LongPollServer struct {
	Error   *Error `json:"error"`
	Server  string `json:"server,omitempty"`
	Key     string `json:"key,omitempty"`
	TS      any    `json:"ts,omitempty"`
	PTS     any    `json:"pts,omitempty"`
	Version int    `json:"version,omitempty"`
}
