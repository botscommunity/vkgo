package types

type SetGroupLongPoll struct {
	GroupID int
	Enabled bool
	Version float32

	Message        bool
	ReplyMessage   bool
	EditMessage    bool
	EnableMessage  bool
	DisableMessage bool
	Typing         bool
	Callback       bool

	Photo               bool
	PhotoComment        bool
	EditPhotoComment    bool
	DeletePhotoComment  bool
	RestorePhotoComment bool

	Audio bool

	Video               bool
	VideoComment        bool
	EditVideoComment    bool
	DeleteVideoComment  bool
	RestoreVideoComment bool

	Post             bool
	Repost           bool
	ReplyPost        bool
	EditReplyPost    bool
	DeleteReplyPost  bool
	RestoreReplyPost bool

	BlockUser   bool
	UnblockUser bool

	Like   bool
	Unlike bool

	BoardPost        bool
	EditBoardPost    bool
	DeleteBoardPost  bool
	RestoreBoardPost bool

	MarketComment        bool
	EditMarketComment    bool
	DeleteMarketComment  bool
	RestoreMarketComment bool

	JoinGroup  bool
	LeaveGroup bool

	Vote bool

	EditOwners          bool
	ChangeGroupSettings bool
	ChangeGroupPhoto    bool

	Donate             bool
	ProlongedDonate    bool
	ExpiredDonate      bool
	CancelDonate       bool
	ChangeDonatePrice  bool
	WithdrawMoney      bool
	WithdrawMoneyError bool
}
