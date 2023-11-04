package responses

type CountersReply struct {
	Error    *Error   `json:"error"`
	Response Counters `json:"response,omitempty"`
}

type Counters struct {
	AppRequests          int              `json:"app_requests,omitempty"`
	Faves                int              `json:"faves,omitempty"`
	DiscoverBadge        int              `json:"menu_discover_badge,omitempty"`
	ClipsBadge           int              `json:"menu_clips_badge,omitempty"`
	SuperAppFriendsBadge int              `json:"menu_superapp_friends_badge,omitempty"`
	NewClipsBadge        int              `json:"menu_new_clips_badge,omitempty"`
	Messages             int              `json:"messages,omitempty"`
	SDK                  int              `json:"sdk,omitempty"`
	Calls                int              `json:"calls,omitempty"`
	MessagesFolders      []MessagesFolder `json:"messages_folders,omitempty"`
}

type MessagesFolder struct {
	ID      int `json:"folder_id,omitempty"`
	Count   int `json:"total_count,omitempty"`
	UnMuted int `json:"unmuted_count,omitempty"`
}
