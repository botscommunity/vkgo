package types

type Message struct {
	GroupID        int
	ChatID         int
	MessageID      int
	MessagesID     []int
	ChatMessageID  int
	ChatMessagesID []int
	Fields         []string
	Extended       bool
	Characters     int
}
