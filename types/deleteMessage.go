package types

type DeleteMessage struct {
	GroupID        int
	ChatID         int
	MessageID      int
	MessagesID     []int
	ChatMessageID  int
	ChatMessagesID []int
	Spam           bool
	Everyone       bool
}
