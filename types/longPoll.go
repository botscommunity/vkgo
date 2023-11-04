package types

type UserLongPollServer struct {
	GroupID int
	Version int
	NeedPTS bool
}

type GroupLongPollServer struct {
	GroupID int
}
