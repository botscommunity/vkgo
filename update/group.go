package update

type JoinGroup struct {
	UserID   int    `json:"user_id,omitempty"`
	JoinType string `json:"join_type,omitempty"`
}

type LeaveGroup struct {
	UserID int `json:"user_id,omitempty"`
	Self   int `json:"self,omitempty"`
}
