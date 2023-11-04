package update

type Action struct {
	Type     string      `json:"type,omitempty"`
	MemberID int         `json:"member_id,omitempty"`
	Text     string      `json:"text,omitempty"`
	EMail    string      `json:"email,omitempty"`
	Photo    PhotoAction `json:"photo,omitempty"`
}

type PhotoAction struct {
	X50  string `json:"photo_50,omitempty"`
	X100 string `json:"photo_100,omitempty"`
	X200 string `json:"photo_200,omitempty"`
}
