package responses

type SendMessages struct {
	Error    *Error        `json:"error"`
	Messages []SendMessage `json:"response,omitempty"`
}

type SendMessage struct {
	Error         *Error `json:"error"`
	UserID        int    `json:"user_id,omitempty"`
	ChatID        int    `json:"peer_id,omitempty"`
	MessageID     int    `json:"message_id,omitempty"`
	ChatMessageID int    `json:"conversation_message_id,omitempty"`
}
