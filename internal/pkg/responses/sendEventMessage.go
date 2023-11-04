package responses

type SendEventMessage struct {
	Error *Error `json:"error"`
	Code  int    `json:"response,omitempty"`
}
