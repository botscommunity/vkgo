package types

type Status struct {
	GroupID int
	UserID  int
}

type SetStatus struct {
	GroupID int
	Text    string
}
