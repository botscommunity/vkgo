package update

type AttachmentDocument struct {
	ID        int    `json:"id,omitempty"`
	OwnerID   int    `json:"owner_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Size      int    `json:"size,omitempty"`
	Extension string `json:"ext,omitempty"`
	Date      int    `json:"date,omitempty"`
	Type      int    `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	IsUnsafe  int    `json:"is_unsafe,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
}

type PreviewDocument struct {
	Photo PhotoDocument `json:"photo,omitempty"`
	Video VideoDocument `json:"video,omitempty"`
}

type PhotoDocument struct {
	Sizes []PhotoDocumentSize `json:"sizes,omitempty"`
}

type PhotoDocumentSize struct {
	Type   string `json:"type,omitempty"`
	Source string `json:"src,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type VideoDocument struct {
	Source   string `json:"src,omitempty"`
	FileSize string `json:"file_size,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}
