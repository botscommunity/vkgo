package update

type AppNotification struct {
	UserID  int    `json:"user_id,omitempty"`
	ID      int    `json:"app_id,omitempty"`
	Payload string `json:"payload,omitempty"`
	GroupID int    `json:"group_id,omitempty"`
}
