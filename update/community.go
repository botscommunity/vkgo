package update

type EditOwners struct {
	AdminID  int `json:"admin_id,omitempty"`
	UserID   int `json:"user_id,omitempty"`
	OldLevel int `json:"level_old,omitempty"`
	NewLevel int `json:"level_new,omitempty"`
}

type ControlGroup struct {
	UserID  int          `json:"user_id,omitempty"`
	Changes GroupChanges `json:"changes,omitempty"`
}

type GroupChanges struct {
	Topics GroupChange `json:"topics,omitempty"`
	Photos GroupChange `json:"photos,omitempty"`
	Videos GroupChange `json:"video,omitempty"`
	Audios GroupChange `json:"audio,omitempty"`
	Links  GroupChange `json:"links,omitempty"`
	Market GroupChange `json:"market,omitempty"`
}

type GroupChange struct {
	Title             string `json:"title,omitempty"`
	Description       string `json:"description,omitempty"`
	Access            any    `json:"access,omitempty"`
	ScreenName        string `json:"screen_name,omitempty"`
	PublicCategory    any    `json:"public_caterogy,omitempty"`
	PublicSubcategory any    `json:"public_subcategory,omitempty"`
	AgeLimits         int    `json:"age_limits,omitempty"`
	Website           string `json:"website,omitempty"`
	EnablePhoto       any    `json:"enable_photo,omitempty"`
	OldValue          int    `json:"old_value,omitempty"`
	NewValue          int    `json:"new_value,omitempty"`
}

type ChangeGroupPhoto struct {
	UserID int             `json:"user_id,omitempty"`
	Photo  AttachmentPhoto `json:"photo,omitempty"`
}
