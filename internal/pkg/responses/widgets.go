package responses

type UpdateWidgets struct {
	Error    *Error `json:"error"`
	Response int    `json:"response,omitempty"`
}
