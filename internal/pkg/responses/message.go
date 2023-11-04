package responses

import (
	"github.com/botscommunity/vkgo/update"
)

type MessagesReply struct {
	Error    *Error   `json:"error"`
	Response Messages `json:"response"`
}

type Messages struct {
	Error    *Error           `json:"error"`
	Count    int              `json:"count,omitempty"`
	Messages []update.Message `json:"items,omitempty"`
	Profiles []User           `json:"profiles,omitempty"`
	Groups   []Group          `json:"groups,omitempty"`
}

func (messages Messages) Decode() (Messages, error) {
	for _, message := range messages.Messages {
		if payload, err := update.DecodePayload(message.RawPayload); err == nil {
			message.Payload = payload
		}
	}

	return messages, nil
}
