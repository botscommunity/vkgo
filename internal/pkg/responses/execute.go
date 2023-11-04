package responses

type Execute struct {
	Error  *Error `json:"error"`
	Result any    `json:"response,omitempty"`
}
