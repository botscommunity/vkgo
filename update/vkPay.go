package update

type VKPayTransaction struct {
	UserID      int    `json:"from_id,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Date        int    `json:"date,omitempty"`
}
