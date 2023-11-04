package responses

type BannedUsersReply struct {
	Error    *Error      `json:"error"`
	Response BannedUsers `json:"response"`
}

type BannedUsers struct {
	Error *Error       `json:"error"`
	Count int          `json:"count,omitempty"`
	Bans  []BannedUser `json:"items,omitempty"`
}

type BannedUser struct {
	Error   *Error      `json:"error"`
	Type    string      `json:"type,omitempty"`
	Info    InfoBanUser `json:"ban_info,omitempty"`
	Profile User        `json:"profile,omitempty"`
}

type InfoBanUser struct {
	AdminID     int    `json:"admin_id,omitempty"`
	Comment     string `json:"comment,omitempty"`
	ShowComment bool   `json:"comment_visible,omitempty"`
	Date        int    `json:"date,omitempty"`
	EndDate     int    `json:"end_date,omitempty"`
	Reason      int    `json:"reason,omitempty"`
}
