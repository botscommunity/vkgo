package scene

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/update"
)

type Func func(*api.Bot, update.Update, Scenes)

type Scenes struct {
	MessageFunc        MessageFunc
	ReplyMessageFunc   ReplyMessageFunc
	EditMessageFunc    EditMessageFunc
	EnableMessageFunc  EnableMessageFunc
	DisableMessageFunc DisableMessageFunc
	TypingMessageFunc  TypingMessageFunc
	CallbackFunc       CallbackFunc

	PhotoFunc               PhotoFunc
	PhotoCommentFunc        PhotoCommentFunc
	EditPhotoCommentFunc    EditPhotoCommentFunc
	RestorePhotoCommentFunc RestorePhotoCommentFunc
	DeletePhotoCommentFunc  DeletePhotoCommentFunc

	NewAudioFunc NewAudioFunc

	NewVideoFunc            NewVideoFunc
	VideoCommentFunc        VideoCommentFunc
	EditVideoCommentFunc    EditVideoCommentFunc
	RestoreVideoCommentFunc RestoreVideoCommentFunc
	DeleteVideoCommentFunc  DeleteVideoCommentFunc

	PostFunc               PostFunc
	RepostFunc             RepostFunc
	PostCommentFunc        PostCommentFunc
	EditPostCommentFunc    EditPostCommentFunc
	RestorePostCommentFunc RestorePostCommentFunc
	DeletePostCommentFunc  DeletePostCommentFunc

	LikeFunc   LikeFunc
	UnlikeFunc UnlikeFunc

	BoardCommentFunc        BoardCommentFunc
	EditBoardCommentFunc    EditBoardCommentFunc
	RestoreBoardCommentFunc RestoreBoardCommentFunc
	DeleteBoardCommentFunc  DeleteBoardCommentFunc

	MarketCommentFunc        MarketCommentFunc
	EditMarketCommentFunc    EditMarketCommentFunc
	RestoreMarketCommentFunc RestoreMarketCommentFunc
	DeleteMarketCommentFunc  DeleteMarketCommentFunc

	OrderMarketFunc     OrderMarketFunc
	EditOrderMarketFunc EditOrderMarketFunc

	JoinGroupFunc  JoinGroupFunc
	LeaveGroupFunc LeaveGroupFunc

	BlockUserFunc   BlockUserFunc
	UnblockUserFunc UnblockUserFunc

	VoteFunc             VoteFunc
	EditOwnersFunc       EditOwnersFunc
	ControlGroupFunc     ControlGroupFunc
	ChangeGroupPhotoFunc ChangeGroupPhotoFunc
	VKPayTransactionFunc VKPayTransactionFunc
	AppNotificationFunc  AppNotificationFunc

	DonateFunc             DonateFunc
	ProlongedDonateFunc    ProlongedDonateFunc
	ExpiredDonateFunc      ExpiredDonateFunc
	CancelDonateFunc       CancelDonateFunc
	ChangeDonatePriceFunc  ChangeDonatePriceFunc
	WithdrawMoneyFunc      WithdrawMoneyFunc
	WithdrawMoneyErrorFunc WithdrawMoneyErrorFunc
}

type MessageFunc func(*api.Bot, update.Message)
type ReplyMessageFunc func(*api.Bot, update.Message)
type EditMessageFunc func(*api.Bot, update.Message)
type EnableMessageFunc func(*api.Bot, update.EnableMessage)
type DisableMessageFunc func(*api.Bot, update.DisableMessage)
type TypingMessageFunc func(*api.Bot, update.Typing)
type CallbackFunc func(*api.Bot, update.Callback)

type PhotoFunc func(*api.Bot, update.Photo)
type PhotoCommentFunc func(*api.Bot, update.PhotoComment)
type EditPhotoCommentFunc PhotoCommentFunc
type RestorePhotoCommentFunc PhotoCommentFunc
type DeletePhotoCommentFunc PhotoCommentFunc

type NewAudioFunc func(*api.Bot, update.Audio)

type NewVideoFunc func(*api.Bot, update.Video)
type VideoCommentFunc func(*api.Bot, update.VideoComment)
type EditVideoCommentFunc VideoCommentFunc
type RestoreVideoCommentFunc VideoCommentFunc
type DeleteVideoCommentFunc func(*api.Bot, update.DeleteVideoComment)

type PostFunc func(*api.Bot, update.Post)
type RepostFunc PostFunc
type PostCommentFunc func(*api.Bot, update.PostComment)
type EditPostCommentFunc PostCommentFunc
type RestorePostCommentFunc PostCommentFunc
type DeletePostCommentFunc func(*api.Bot, update.DeletePostComment)

type LikeFunc func(*api.Bot, update.Like)
type UnlikeFunc func(*api.Bot, update.Unlike)

type BoardCommentFunc func(*api.Bot, update.BoardComment)
type EditBoardCommentFunc BoardCommentFunc
type RestoreBoardCommentFunc BoardCommentFunc
type DeleteBoardCommentFunc func(*api.Bot, update.DeleteBoardComment)

type MarketCommentFunc func(*api.Bot, update.MarketComment)
type EditMarketCommentFunc MarketCommentFunc
type RestoreMarketCommentFunc MarketCommentFunc
type DeleteMarketCommentFunc func(*api.Bot, update.DeleteMarketComment)

type OrderMarketFunc func(*api.Bot, update.OrderMarket)
type EditOrderMarketFunc OrderMarketFunc

type JoinGroupFunc func(*api.Bot, update.JoinGroup)
type LeaveGroupFunc func(*api.Bot, update.LeaveGroup)

type BlockUserFunc func(*api.Bot, update.BlockUser)
type UnblockUserFunc func(*api.Bot, update.UnblockUser)

type VoteFunc func(*api.Bot, update.Vote)
type EditOwnersFunc func(*api.Bot, update.EditOwners)
type ControlGroupFunc func(*api.Bot, update.ControlGroup)
type ChangeGroupPhotoFunc func(*api.Bot, update.ChangeGroupPhoto)
type VKPayTransactionFunc func(*api.Bot, update.VKPayTransaction)
type AppNotificationFunc func(*api.Bot, update.AppNotification)

type DonateFunc func(*api.Bot, update.ProlongedDonate)
type ProlongedDonateFunc func(*api.Bot, update.ProlongedDonate)
type ExpiredDonateFunc func(*api.Bot, update.ExpiredDonate)
type CancelDonateFunc func(*api.Bot, update.ExpiredDonate)
type ChangeDonatePriceFunc func(*api.Bot, update.ChangeDonatePrice)
type WithdrawMoneyFunc func(*api.Bot, update.WithdrawMoney)
type WithdrawMoneyErrorFunc func(*api.Bot, update.WithdrawMoneyError)

func New() *Scenes {
	return new(Scenes)
}
