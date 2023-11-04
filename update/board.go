package update

type BoardComment struct {
	ID          int          `json:"id,omitempty"`
	UserID      int          `json:"from_id,omitempty"`
	Date        int          `json:"date,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Likes       Likes        `json:"likes,omitempty"`

	TopicID      int `json:"topic_id,omitempty"`
	TopicOwnerID int `json:"topic_owner_id,omitempty"`
}

type DeleteBoardComment struct {
	TopicOwnerID int `json:"topic_owner_id,omitempty"`
	TopicID      int `json:"topic_id,omitempty"`
	ID           int `json:"id,omitempty"`
}
