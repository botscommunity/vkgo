package update

type ProlongedDonate struct {
	Amount           int     `json:"amount,omitempty"`
	AmountWithoutFee float32 `json:"amount_without_fee,omitempty"`
	UserID           int     `json:"user_id,omitempty"`
}

type ExpiredDonate struct {
	UserID int `json:"user_id,omitempty"`
}

type ChangeDonatePrice struct {
	OldAmount            int     `json:"amount_old,omitempty"`
	NewAmount            int     `json:"amount_new,omitempty"`
	AmountDiff           float32 `json:"amount_diff,omitempty"`
	AmountDiffWithoutFee float32 `json:"amount_diff_without_fee,omitempty"`
	UserID               int     `json:"user_id,omitempty"`
}

type WithdrawMoney struct {
	Amount           float32 `json:"amount,omitempty"`
	AmountWithoutFee float32 `json:"amount_without_fee,omitempty"`
}

type WithdrawMoneyError struct {
	Reason string `json:"reason,omitempty"`
}
