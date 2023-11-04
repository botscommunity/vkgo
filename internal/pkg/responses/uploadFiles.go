package responses

type UploadPhotoFile struct {
	Server int    `json:"server,omitempty"`
	Photo  string `json:"photo,omitempty"`
	Hash   string `json:"hash,omitempty"`
}

type UploadDocumentFile struct {
	File string `json:"file,omitempty"`
}
