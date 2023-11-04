package responses

type OnlineStatus struct {
	Error *Error `json:"error"`
	Code  int    `json:"response,omitempty"`
}
