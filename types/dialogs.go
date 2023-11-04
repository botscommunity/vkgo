package types

type Dialogs struct {
	Offset         int
	Count          int
	StartMessageID int
	PreviewLength  int
	Unread         bool
	Important      bool
	Unanswered     bool
	UserID         string
}
