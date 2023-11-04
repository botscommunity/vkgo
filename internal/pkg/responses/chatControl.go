package responses

import (
	"github.com/botscommunity/vkgo/update"
)

type CreateChatReply struct {
	Error    *Error     `json:"error"`
	Response CreateChat ` json:"response"`
}

type CreateChat struct {
	Error  *Error `json:"error"`
	ChatID int    `json:"chat_id,omitempty"`
	Users  []int  `json:"peer_ids,omitempty"`
}

type ChatReply struct {
	Error    *Error ` json:"error"`
	Response Chats  ` json:"response,omitempty"`
}

type Chats struct {
	Error  *Error  `json:"error"`
	Count  int     `json:"count,omitempty"`
	Chats  []Chat  `json:"items,omitempty"`
	Users  []User  `json:"profiles,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

type Chat struct {
	Error *Error `json:"error"`
	Chat  struct {
		ID      int    `json:"id,omitempty"`
		Type    string `json:"type,omitempty"`
		LocalID int    `json:"local_id,omitempty"`
	} `json:"peer,omitempty"`
	InRead               int          `json:"in_read,omitempty"`
	OutRead              int          `json:"out_read,omitempty"`
	UnreadCount          int          `json:"unread_count,omitempty"`
	Important            bool         `json:"important,omitempty"`
	Unanswered           bool         `json:"unanswered,omitempty"`
	PushSettings         PushSettings `json:"push_settings,omitempty"`
	CanWrite             CanWriteChat `json:"can_write,omitempty"`
	ChatSettings         ChatSettings `json:"chat_settings,omitempty"`
	SortID               ChatSortID   `json:"sort_id,omitempty"`
	LastMessageID        int          `json:"last_message_id,omitempty"`
	LastChatMessageID    int          `json:"last_conversation_message_id,omitempty"`
	InReadChatMessageID  int          `json:"in_read_cmid,omitempty"`
	OutReadChatMessageID int          `json:"out_read_cmid,omitempty"`
	MarkedUnread         bool         `json:"is_marked_unread,omitempty"`
	CanSendMoney         bool         `json:"can_send_money,omitempty"`
	CanReceiveMoney      bool         `json:"can_receive_money,omitempty"`
}

type PushSettings struct {
	UntilDisabled   bool `json:"disabled_until,omitempty"`
	ForeverDisabled bool `json:"disabled_forever,omitempty"`
	Sound           bool `json:"sound,omitempty"`
	NoSound         bool `json:"no_sound,omitempty"`
}

type ChatSortID struct {
	MajorID int `json:"major_id,omitempty"`
	MinorID int `json:"minor_id,omitempty"`
}

type CanWriteChat struct {
	Allowed bool `json:"allowed,omitempty"`
	Reason  int  `json:"reason,omitempty"`
}

type ChatSettings struct {
	MembersCount   int                `json:"members_count,omitempty"`
	Title          string             `json:"title,omitempty"`
	PinnedMessage  PinMessage         `json:"pinned_message,omitempty"`
	State          string             `json:"state,omitempty"`
	Photo          update.PhotoAction `json:"photo,omitempty"`
	ActivesID      []int              `json:"active_ids,omitempty"`
	IsGroupChannel bool               `json:"is_group_channel,omitempty"`
}

type ChatLinkReply struct {
	Error    *Error   `json:"error"`
	Response ChatLink ` json:"response,omitempty"`
}

type ChatLink struct {
	Error *Error `json:"error"`
	Link  string `json:"link,omitempty"`
}

type ChatMembersReply struct {
	Error    *Error      `json:"error"`
	Response ChatMembers `json:"response,omitempty"`
}

type ChatMembers struct {
	Error    *Error           `json:"error"`
	Count    int              `json:"count,omitempty"`
	Settings ChatRestrictions `json:"chat_restrictions,omitempty"`
	Members  []ChatMember     `json:"items,omitempty"`
	Users    []User           `json:"profiles,omitempty"`
	Groups   []Group          `json:"groups,omitempty"`
}

type ChatMember struct {
	ID        int  `json:"member_id,omitempty"`
	InvitedBy int  `json:"invited_by,omitempty"`
	IsOwner   bool `json:"is_owner,omitempty"`
	IsAdmin   bool `json:"is_admin,omitempty"`
	CanKick   bool `json:"can_kick,omitempty"`
	JoinDate  int  `json:"join_date,omitempty"`
}

type ChatRestrictions struct {
	AdminsPromoteUsers bool `json:"admins_promote_users,omitempty"`
	OnlyAdminsEditInfo bool `json:"only_admins_edit_info,omitempty"`
	OnlyAdminsEditPin  bool `json:"only_admins_edit_pin,omitempty"`
	OnlyAdminsInvite   bool `json:"only_admins_invite,omitempty"`
	OnlyAdminsKick     bool `json:"only_admins_kick,omitempty"`
}
