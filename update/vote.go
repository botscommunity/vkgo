package update

type Vote struct {
	OwnerID  int `json:"owner_id,omitempty"`
	PollID   int `json:"poll_id,omitempty"`
	OptionID int `json:"option_id,omitempty"`
	UserID   int `json:"user_id,omitempty"`
}
