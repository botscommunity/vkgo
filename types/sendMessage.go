package types

import "time"

type SendMessage struct {
	GroupID         int
	RandomID        int
	Domain          string
	ChatID          int
	ChatsID         []int
	UserID          int
	UsersID         []int
	Text            string
	Attachment      string
	Attachments     []string
	StickerID       int
	Payload         string
	Forward         Forward
	Template        string
	Keyboard        string
	DontParseLinks  bool
	DisableMentions bool
	Latitude        string
	Longitude       string
	Intent          string
	SubscribeID     int
	ExpireTime      int
}

type Forward struct {
	ChatID        int  `json:"peer_id"`
	ChatMessageID int  `json:"conversation_message_ids"`
	IsReply       bool `json:"is_reply"`
}

func GetForward(chatID, chatMessageID int, isReply bool) Forward {
	return Forward{chatID, chatMessageID, isReply}
}

func Random() int {
	return int(time.Now().UnixNano())
}
