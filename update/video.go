package update

type AttachmentVideo Video

type Video struct {
	ID            int          `json:"id,omitempty"`
	OwnerID       int          `json:"owner_id,omitempty"`
	Title         string       `json:"title,omitempty"`
	Description   string       `json:"description,omitempty"`
	Duration      int          `json:"duration,omitempty"`
	Photos        []VideoPhoto `json:"image,omitempty"`
	FirstFrame    []VideoPhoto `json:"first_frame,omitempty"`
	Date          int          `json:"date,omitempty"`
	AddingDate    int          `json:"adding_date,omitempty"`
	Views         int          `json:"views,omitempty"`
	LocalViews    int          `json:"local_views,omitempty"`
	Comments      int          `json:"comments,omitempty"`
	Player        string       `json:"player,omitempty"`
	Platform      string       `json:"platform,omitempty"`
	CanAdd        int          `json:"can_add,omitempty"`
	IsPrivate     int          `json:"is_private,omitempty"`
	AccessKey     string       `json:"access_key,omitempty"`
	Processing    int          `json:"processing,omitempty"`
	IsFavorite    bool         `json:"is_favorite,omitempty"`
	CanComment    int          `json:"can_comment,omitempty"`
	CanEdit       int          `json:"can_edit,omitempty"`
	CanLike       int          `json:"can_like,omitempty"`
	CanRepost     int          `json:"can_repost,omitempty"`
	CanSubscribe  int          `json:"can_subscribe,omitempty"`
	CanAddToFaves int          `json:"can_add_to_faves,omitempty"`
	CanAttachLink int          `json:"can_attach_link,omitempty"`
	Width         int          `json:"width,omitempty"`
	Height        int          `json:"height,omitempty"`
	UserID        int          `json:"user_id,omitempty"`
	Converting    int          `json:"converting,omitempty"`
	Added         int          `json:"added,omitempty"`
	Subscribed    int          `json:"is_subscribed,omitempty"`
	Repeat        int          `json:"repeat,omitempty"`
	Type          string       `json:"type,omitempty"`
	Balance       int          `json:"balance,omitempty"`
	LiveStatus    string       `json:"live_status,omitempty"`
	Live          int          `json:"live,omitempty"`
	Upcoming      int          `json:"upcoming,omitempty"`
	Spectators    int          `json:"spectators,omitempty"`
	Likes         Likes        `json:"likes,omitempty"`
	Reposts       Reposts      `json:"reposts,omitempty"`
}

type VideoPhoto struct {
	URL         string `json:"url,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	WithPadding int    `json:"with_padding,omitempty"`
}

type VideoComment struct {
	PostComment

	VideoID      int `json:"video_id,omitempty"`
	VideoOwnerID int `json:"video_owner_id"`
}

type DeleteVideoComment struct {
	OwnerID   int `json:"owner_id,omitempty"`
	ID        int `json:"id,omitempty"`
	UserID    int `json:"user_id,omitempty"`
	DeleterID int `json:"deleter_id,omitempty"`
	VideoID   int `json:"ideo_id,omitempty"`
}
