package responses

type SaveUploadMessagesPhotoReply struct {
	Error  *Error                    `json:"error"`
	Photos []SaveUploadMessagesPhoto `json:"response"`
}

type SaveUploadMessagesPhoto struct {
	Date      int    `json:"date,omitempty"`
	ID        int    `json:"id,omitempty"`
	AlbumID   int    `json:"album_id,omitempty"`
	OwnerID   int    `json:"owner_id,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
}
