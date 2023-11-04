package types

type CreateChat struct {
	GroupID int
	UserID  int
	UsersID []int
	Name    string
	Title   string
}

type Chat struct {
	GroupID  int
	ChatID   int
	ChatsID  []int
	Extended bool
	Fields   []string
}

type ChatLink struct {
	GroupID int
	ChatID  int
	Reset   bool
}

type ChatMembers struct {
	GroupID  int
	ChatID   int
	Offset   int
	Count    int
	Extended bool
	Fields   []string
}
