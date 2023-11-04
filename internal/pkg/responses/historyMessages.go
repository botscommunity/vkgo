package responses

import "github.com/botscommunity/vkgo/update"

type HistoryMessagesReply struct {
	Error    *Error          `json:"error"`
	Response HistoryMessages `json:"response"`
}

type HistoryMessages struct {
	Error    *Error           `json:"error"`
	Count    int              `json:"count,omitempty"`
	Messages []update.Message `json:"items,omitempty"`
	Users    []User           `json:"profiles,omitempty"`
	Chats    []Chat           `json:"conversations,omitempty"`
	Contacts []Contact        `json:"contacts,omitempty"`
}
