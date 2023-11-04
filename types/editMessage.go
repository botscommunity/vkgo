package types

type EditMessage struct {
	GroupID             int
	ChatID              int
	Text                string
	MessageID           int
	ChatMessageID       int
	Attachment          string
	Attachments         []string
	Template            string
	Keyboard            string
	KeepForwardMessages bool
	KeepSnippets        bool
	DontParseLinks      bool
	DisableMentions     bool
	Latitude            string
	Longitude           string
}
