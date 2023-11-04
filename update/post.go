package update

type Post struct {
	ID           int          `json:"id,omitempty"`
	OwnerID      int          `json:"owner_id,omitempty"`
	UserID       int          `json:"from_id,omitempty"`
	CreatedBy    int          `json:"created_by,omitempty"`
	Date         int          `json:"date,omitempty"`
	Text         string       `json:"text,omitempty"`
	ReplyOwnerID int          `json:"reply_owner_id,omitempty"`
	ReplyPostID  int          `json:"reply_post_id,omitempty"`
	FriendsOnly  int          `json:"friends_only,omitempty"`
	Comments     Comments     `json:"comments,omitempty"`
	Copyright    Copyright    `json:"copyright,omitempty"`
	Likes        Likes        `json:"likes,omitempty"`
	Reposts      Reposts      `json:"reposts,omitempty"`
	Views        Views        `json:"views,omitempty"`
	Type         string       `json:"post_type,omitempty"`
	Source       string       `json:"post_source,omitempty"`
	Attachments  []Attachment `json:"attachments,omitempty"`
	Geolocation  Geolocation  `json:"geo,omitempty"`
	SignerID     int          `json:"signer_id,omitempty"`
	CopyHistory  []any        `json:"copy_history,omitempty"`
	CanPin       int          `json:"can_pin,omitempty"`
	CanDelete    int          `json:"can_delete,omitempty"`
	CanEdit      int          `json:"can_edit,omitempty"`
	Pinned       int          `json:"is_pinned,omitempty"`
	MarkedAsAds  int          `json:"marked_as_ads,omitempty"`
	IsFavorite   bool         `json:"is_favorite,omitempty"`
	Donut        struct {
		IsDonut            bool   `json:"is_donut,omitempty"`
		PaidDuration       int    `json:"paid_duration,omitempty"`
		Placeholder        any    `json:"placeholder,omitempty"`
		CanPublishFreeCopy bool   `json:"can_publish_free_copy,omitempty"`
		EditMode           string `json:"edit_mode,omitempty"`
	} `json:"donut,omitempty"`
	PostponedID int `json:"postponed_id,omitempty"`
}

type PostComment struct {
	ID     int    `json:"id,omitempty"`
	UserID int    `json:"from_id,omitempty"`
	Date   int    `json:"date,omitempty"`
	Text   string `json:"text,omitempty"`
	Donut  struct {
		IsDon       bool   `json:"is_don,omitempty"`
		Placeholder string `json:"placeholder,omitempty"`
	} `json:"donut,omitempty"`
	ReplyToUser    int          `json:"reply_to_user,omitempty"`
	ReplyToComment int          `json:"reply_to_comment,omitempty"`
	Attachments    []Attachment `json:"attachments,omitempty"`
	ParentsStack   []int        `json:"parents_stack,omitempty"`
	Thread         Thread       `json:"thread,omitempty"`

	PostID      int `json:"post_id,omitempty"`
	PostOwnerID int `json:"post_owner_id,omitempty"`
}

type DeletePostComment struct {
	OwnerID   int `json:"owner_id,omitempty"`
	ID        int `json:"id,omitempty"`
	DeleterID int `json:"deleter_id,omitempty"`
	PostID    int `json:"post_id,omitempty"`
}
