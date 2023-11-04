package types

type BanUser struct {
	GroupID int
	UserID  int
	Time    int
	Reason  int
	Comment BanUserComment
}

type BanUserComment struct {
	Text string
	Show bool
}

type BannedUsers struct {
	GroupID int
	UserID  int
	Offset  int
	Count   int
	Fields  []string
}
