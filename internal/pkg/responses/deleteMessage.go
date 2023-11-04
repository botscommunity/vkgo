package responses

type DeleteMessages struct {
	Error    *Error          `json:"error"`
	Messages []DeleteMessage `json:"response,omitempty"`
}

type DeleteMessage struct {
	Error         *Error `json:"error"`
	ChatID        int    `json:"peer_id,omitempty"`
	MessageID     int    `json:"message_id,omitempty"`
	ChatMessageID int    `json:"conversation_message_id,omitempty"`
	Response      int    `json:"response,omitempty"`
}

type DeleteMessageError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"description,omitempty"`
}
