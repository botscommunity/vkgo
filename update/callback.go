package update

type Callback struct {
	ChatID        int        `json:"peer_id,omitempty"`
	UserID        int        `json:"user_id,omitempty"`
	EventID       string     `json:"event_id,omitempty"`
	RawPayload    rawPayload `json:"payload,omitempty"`
	Payload       map[string]string
	ChatMessageID int `json:"conversation_message_id,omitempty"`
}
