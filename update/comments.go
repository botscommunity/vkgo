package update

type Comments struct {
	Count         int  `json:"count,omitempty"`
	CanPost       int  `json:"can_post,omitempty"`
	GroupsCanPost bool `json:"groups_can_post,omitempty"`
	CanClose      bool `json:"can_close,omitempty"`
	CanOpen       bool `json:"can_open,omitempty"`
}
