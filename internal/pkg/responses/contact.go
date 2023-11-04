package responses

type Contact struct {
	ID             int    `json:"id,omitempty"`
	UserID         int    `json:"user_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"desc,omitempty"`
	Phone          string `json:"phone,omitempty"`
	EMail          string `json:"email,omitempty"`
	CanWrite       bool   `json:"can_write,omitempty"`
	LastSeenStatus int    `json:"last_seen_status,omitempty"`
	Photo50        int    `json:"photo_50,omitempty"`
	CallsID        int    `json:"calls_id,omitempty"`
}
