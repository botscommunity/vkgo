package responses

type SetGroupLongPoll struct {
	Error *Error `json:"error"`
	Code  int    `json:"response,omitempty"`
}
