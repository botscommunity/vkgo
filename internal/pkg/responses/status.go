package responses

type StatusReply struct {
	Error    *Error `json:"error"`
	Response Status `json:"response"`
}

type Status struct {
	Error *Error `json:"error"`
	Text  string `json:"text,omitempty"`
}

type SetStatus struct {
	Error *Error `json:"error"`
	Code  int    `json:"response,omitempty"`
}
