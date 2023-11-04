package update

type AttachmentPhoto Photo

type Photo struct {
	ID         int         `json:"id,omitempty"`
	AlbumID    int         `json:"album_id,omitempty"`
	OwnerID    int         `json:"owner_id,omitempty"`
	UserID     int         `json:"user_id,omitempty"`
	Date       int         `json:"date,omitempty"`
	Text       string      `json:"text,omitempty"`
	Sizes      []PhotoSize `json:"sizes,omitempty"`
	Width      int         `json:"width,omitempty"`
	Height     int         `json:"height,omitempty"`
	SquareCrop string      `json:"square_crop,omitempty"`
	HasTags    bool        `json:"has_tags,omitempty"`
}

type PhotoSize struct {
	Type   string `json:"type,omitempty"`
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type PhotoComment struct {
	PostComment

	PhotoID      int `json:"photo_id,omitempty"`
	PhotoOwnerID int `json:"photo_owner_id,omitempty"`
}
