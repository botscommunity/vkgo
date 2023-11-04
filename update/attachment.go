package update

type Attachment struct {
	Type     string             `json:"type,omitempty"`
	Photo    AttachmentPhoto    `json:"photo,omitempty"`
	Document AttachmentDocument `json:"doc,omitempty"`
	Video    AttachmentVideo    `json:"video,omitempty"`
	Audio    AttachmentAudio    `json:"audio,omitempty"`
	Market   AttachmentMarket   `json:"market,omitempty"`
}
