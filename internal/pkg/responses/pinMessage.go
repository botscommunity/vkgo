package responses

import (
	"github.com/botscommunity/vkgo/update"
)

type PinMessageReply struct {
	Error    *Error     `json:"error"`
	Response PinMessage `json:"response"`
}

type PinMessage struct {
	Error       *Error              `json:"error"`
	ID          int                 `json:"id,omitempty"`
	Date        int                 `json:"date,omitempty"`
	UserID      int                 `json:"from_id,omitempty"`
	Text        string              `json:"text,omitempty"`
	Attachments []update.Attachment `json:"attachments,omitempty"`
	Geolocation update.Geolocation  `json:"geo,omitempty"`
	Forwards    []string            `json:"fwd_messages,omitempty"`
}
