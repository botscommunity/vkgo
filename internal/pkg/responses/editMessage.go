package responses

type EditMessage struct {
	Error *Error `json:"error"`
	Code  int    `json:"response,omitempty"`
}
