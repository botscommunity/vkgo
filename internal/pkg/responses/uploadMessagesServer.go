package responses

type UploadMessagesPhotoServerReply struct {
	Error    *Error                    `json:"error"`
	Response UploadMessagesPhotoServer `json:"response,omitempty"`
}

type UploadMessagesPhotoServer struct {
	Error     *Error `json:"error"`
	AlbumID   int    `json:"album_id,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	UploadURL string `json:"upload_url,omitempty"`
}

type UploadMessagesDocumentServerReply struct {
	Error    *Error                       `json:"error"`
	Response UploadMessagesDocumentServer `json:"response,omitempty"`
}

type UploadMessagesDocumentServer struct {
	Error     *Error `json:"error"`
	UploadURL string `json:"upload_url,omitempty"`
}
