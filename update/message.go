package update

import (
	"github.com/botscommunity/vkgo/keyboard"
)

type Message struct {
	ID             int          `json:"id,omitempty"`
	Date           int          `json:"date,omitempty"`
	ChatID         int          `json:"peer_id,omitempty"`
	UserID         int          `json:"from_id,omitempty"`
	Text           string       `json:"text,omitempty"`
	RandomID       int          `json:"random_id,omitempty"`
	Referral       string       `json:"ref,omitempty"`
	ReferralSource string       `json:"ref_source,omitempty"`
	Attachments    []Attachment `json:"attachments,omitempty"`
	Important      bool         `json:"important,omitempty"`
	Geolocation    Geolocation  `json:"geo,omitempty"`
	RawPayload     rawPayload   `json:"payload,omitempty"`
	Payload        map[string]string
	Keyboard       keyboard.Keyboard `json:"keyboard,omitempty"`
	Forwards       []*Message        `json:"fwd_messages,omitempty"`
	Reply          *Message          `json:"reply_message,omitempty"`
	Action         Action            `json:"action,omitempty"`
	AuthorID       int               `json:"admin_author_id,omitempty"`
	ChatMessageID  int               `json:"conversation_message_id,omitempty"`
	Cropped        bool              `json:"is_cropped,omitempty"`
	MembersCount   int               `json:"members_count,omitempty"`
	UpdateTime     int               `json:"update_time,omitempty"`
	Listened       bool              `json:"was_listened,omitempty"`
	PinnedAt       int               `json:"pinned_at,omitempty"`
	Tag            string            `json:"message_tag,omitempty"`
	Out            int               `json:"out,omitempty"`
	Hidden         bool              `json:"is_hidden,omitempty"`
	ExpireTime     int               `json:"expire_ttl"`
}

type EnableMessage struct {
	UserID int    `json:"user_id,omitempty"`
	Key    string `json:"key,omitempty"`
}

type DisableMessage struct {
	UserID int `json:"user_id,omitempty"`
}

type Typing struct {
	State  string `json:"state,omitempty"`
	UserID int    `json:"from_id,omitempty"`
	ToID   int    `json:"to_id,omitempty"`
}

type forwardsMain func(*Message, int)

func (message Message) ForEachForwards(execute forwardsMain) {
	for index, message := range message.Forwards {
		execute(message, index)
	}

	if message.Reply != nil {
		execute(message.Reply, len(message.Forwards))
	}
}
