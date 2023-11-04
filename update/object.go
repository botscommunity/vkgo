package update

type Object struct {
	Message    Message    `json:"message,omitempty"`
	ClientInfo ClientInfo `json:"client_info,omitempty"`
}
