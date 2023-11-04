package update

type Reposts struct {
	Count         int `json:"count,omitempty"`
	Posts         int `json:"wall_count,omitempty"`
	MessagesCount int `json:"mail_count,omitempty"`
	Reposted      int `json:"user_reposted,omitempty"`
}
