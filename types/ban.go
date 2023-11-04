package types

type Ban struct {
	OwnerID int
}

type Banned struct {
	Offset int
	Count  int
}

type UnbanUser struct {
	GroupID int
	UserID  int
}
