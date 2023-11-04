package update

type BlockUser struct {
	AdminID     int    `json:"admin_id,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	UnblockDate int    `json:"unblock_date,omitempty"`
	Reason      int    `json:"reason,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type UnblockUser struct {
	AdminID int `json:"admin_id,omitempty"`
	UserID  int `json:"user_id,omitempty"`
	EndDate int `json:"by_end_date,omitempty"`
}
