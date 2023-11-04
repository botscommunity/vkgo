package types

type HistoryMessages struct {
	GroupID        int
	ChatID         int
	UserID         int
	StartMessageID int
	Offset         int
	Count          int
	Rev            bool
	Extended       bool
	Fields         []string
}
