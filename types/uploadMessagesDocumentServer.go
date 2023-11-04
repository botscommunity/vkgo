package types

type UploadMessagesPhotoServer struct {
	ChatID int
}

type UploadMessagesDocumentServer struct {
	ChatID int
	Type   string
}
