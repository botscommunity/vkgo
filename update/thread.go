package update

type Thread struct {
	Count           int   `json:"count,omitempty"`
	Items           []any `json:"items,omitempty"`
	CanPost         bool  `json:"can_post,omitempty"`
	ShowReplyButton bool  `json:"show_reply_button,omitempty"`
	GroupsCanPost   bool  `json:"groups_can_post,omitempty"`
}
