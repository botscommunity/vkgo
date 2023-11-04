package scene

import (
	"encoding/json"
	"fmt"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/update"
)

func newMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	object := update.Object{}
	if err := json.Unmarshal(thisUpdate.Object, &object); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("NewMessage error is %s", err.Error()))
		return
	}

	object.Message.Payload, _ = update.DecodePayload(object.Message.RawPayload)

	if scenes.MessageFunc != nil {
		scenes.MessageFunc(bot, object.Message)
	}
}

func replyMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	object := update.Message{}
	if err := json.Unmarshal(thisUpdate.Object, &object); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageReply error is %s", err.Error()))
		return
	}

	object.Payload, _ = update.DecodePayload(object.RawPayload)

	if scenes.ReplyMessageFunc != nil {
		scenes.ReplyMessageFunc(bot, object)
	}
}

func editMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	object := update.Message{}
	if err := json.Unmarshal(thisUpdate.Object, &object); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageEdit error is %s", err.Error()))
		return
	}

	object.Payload, _ = update.DecodePayload(object.RawPayload)

	if scenes.EditMessageFunc != nil {
		scenes.EditMessageFunc(bot, object)
	}
}

func enableMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	enableMessage := update.EnableMessage{}
	if err := json.Unmarshal(thisUpdate.Object, &enableMessage); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageAllow error is %s", err.Error()))
		return
	}

	if scenes.EnableMessageFunc != nil {
		scenes.EnableMessageFunc(bot, enableMessage)
	}
}

func disableMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	disableMessage := update.DisableMessage{}
	if err := json.Unmarshal(thisUpdate.Object, &disableMessage); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageDeny error is %s", err.Error()))
		return
	}

	if scenes.DisableMessageFunc != nil {
		scenes.DisableMessageFunc(bot, disableMessage)
	}
}

func typingMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	typing := update.Typing{}
	if err := json.Unmarshal(thisUpdate.Object, &typing); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageTyping error is %s", err.Error()))
		return
	}

	if scenes.TypingMessageFunc != nil {
		scenes.TypingMessageFunc(bot, typing)
	}
}

func eventMessage(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	callback := update.Callback{}
	if err := json.Unmarshal(thisUpdate.Object, &callback); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MessageEvent error is %s", err.Error()))
		return
	}

	callback.Payload, _ = update.DecodePayload(callback.RawPayload)

	if scenes.CallbackFunc != nil {
		scenes.CallbackFunc(bot, callback)
	}
}

func newPhoto(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	photo := update.Photo{}
	if err := json.Unmarshal(thisUpdate.Object, &photo); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PhotoNew error is %s", err.Error()))
		return
	}

	if scenes.PhotoFunc != nil {
		scenes.PhotoFunc(bot, photo)
	}
}

func newPhotoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PhotoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PhotoCommentNew error is %s", err.Error()))
		return
	}

	if scenes.PhotoCommentFunc != nil {
		scenes.PhotoCommentFunc(bot, comment)
	}
}

func editPhotoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PhotoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PhotoCommentEdit error is %s", err.Error()))
		return
	}

	if scenes.EditPhotoCommentFunc != nil {
		scenes.EditPhotoCommentFunc(bot, comment)
	}
}

func restorePhotoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PhotoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PhotoCommentRestore error is %s", err.Error()))
		return
	}

	if scenes.RestorePhotoCommentFunc != nil {
		scenes.RestorePhotoCommentFunc(bot, comment)
	}
}

func deletePhotoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PhotoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PhotoCommentDelete error is %s", err.Error()))
		return
	}

	if scenes.DeletePhotoCommentFunc != nil {
		scenes.DeletePhotoCommentFunc(bot, comment)
	}
}

func newAudio(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	audio := update.Audio{}
	if err := json.Unmarshal(thisUpdate.Object, &audio); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("AudioNew error is %s", err.Error()))
		return
	}

	if scenes.NewAudioFunc != nil {
		scenes.NewAudioFunc(bot, audio)
	}
}

func newVideo(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	video := update.Video{}
	if err := json.Unmarshal(thisUpdate.Object, &video); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VideoNew error is %s", err.Error()))
		return
	}

	if scenes.NewVideoFunc != nil {
		scenes.NewVideoFunc(bot, video)
	}
}

func newVideoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.VideoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VideoCommentNew error is %s", err.Error()))
		return
	}

	if scenes.VideoCommentFunc != nil {
		scenes.VideoCommentFunc(bot, comment)
	}
}

func editVideoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.VideoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VideoCommentEdit error is %s", err.Error()))
		return
	}

	if scenes.EditVideoCommentFunc != nil {
		scenes.EditVideoCommentFunc(bot, comment)
	}
}

func restoreVideoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.VideoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VideoCommentRestore error is %s", err.Error()))
		return
	}

	if scenes.RestoreVideoCommentFunc != nil {
		scenes.RestoreVideoCommentFunc(bot, comment)
	}
}

func deleteVideoComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.DeleteVideoComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VideoCommentDelete error is %s", err.Error()))
		return
	}

	if scenes.DeleteVideoCommentFunc != nil {
		scenes.DeleteVideoCommentFunc(bot, comment)
	}
}

func newWallPost(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	post := update.Post{}
	if err := json.Unmarshal(thisUpdate.Object, &post); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallPostNew error is %s", err.Error()))
		return
	}

	if scenes.PostFunc != nil {
		scenes.PostFunc(bot, post)
	}
}

func repostWall(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	post := update.Post{}
	if err := json.Unmarshal(thisUpdate.Object, &post); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallRepost error is %s", err.Error()))
		return
	}

	if scenes.RepostFunc != nil {
		scenes.RepostFunc(bot, post)
	}
}

func newWallReply(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PostComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallReplyNew error is %s", err.Error()))
		return
	}

	if scenes.PostCommentFunc != nil {
		scenes.PostCommentFunc(bot, comment)
	}
}

func editWallReply(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PostComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallReplyEdit error is %s", err.Error()))
		return
	}

	if scenes.EditPostCommentFunc != nil {
		scenes.EditPostCommentFunc(bot, comment)
	}
}

func restoreWallReply(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.PostComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallReplyRestore error is %s", err.Error()))
		return
	}

	if scenes.RestorePostCommentFunc != nil {
		scenes.RestorePostCommentFunc(bot, comment)
	}
}

func deleteWallReply(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.DeletePostComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("WallReplyDelete error is %s", err.Error()))
		return
	}

	if scenes.DeletePostCommentFunc != nil {
		scenes.DeletePostCommentFunc(bot, comment)
	}
}

func like(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	like := update.Like{}
	if err := json.Unmarshal(thisUpdate.Object, &like); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("LikeAdd error is %s", err.Error()))
		return
	}

	if scenes.LikeFunc != nil {
		scenes.LikeFunc(bot, like)
	}
}

func unlike(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	unlike := update.Unlike{}
	if err := json.Unmarshal(thisUpdate.Object, &unlike); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("LikeRemove error is %s", err.Error()))
		return
	}

	if scenes.UnlikeFunc != nil {
		scenes.UnlikeFunc(bot, unlike)
	}
}

func newBoardPost(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.BoardComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("BoardPostNew error is %s", err.Error()))
		return
	}

	if scenes.BoardCommentFunc != nil {
		scenes.BoardCommentFunc(bot, comment)
	}
}

func editBoardPost(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.BoardComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("BoardPostEdit error is %s", err.Error()))
		return
	}

	if scenes.EditBoardCommentFunc != nil {
		scenes.EditBoardCommentFunc(bot, comment)
	}
}

func restoreBoardPost(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.BoardComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("BoardPostRestore error is %s", err.Error()))
		return
	}

	if scenes.RestoreBoardCommentFunc != nil {
		scenes.RestoreBoardCommentFunc(bot, comment)
	}
}

func deleteBoardPost(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.DeleteBoardComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("BoardPostDelete error is %s", err.Error()))
		return
	}

	if scenes.DeleteBoardCommentFunc != nil {
		scenes.DeleteBoardCommentFunc(bot, comment)
	}
}

func marketComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.MarketComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketComment error is %s", err.Error()))
		return
	}

	if scenes.MarketCommentFunc != nil {
		scenes.MarketCommentFunc(bot, comment)
	}
}

func editMarketComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.MarketComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketCommentEdit error is %s", err.Error()))
		return
	}

	if scenes.EditMarketCommentFunc != nil {
		scenes.EditMarketCommentFunc(bot, comment)
	}
}

func restoreMarketComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.MarketComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketCommentRestore error is %s", err.Error()))
		return
	}

	if scenes.RestoreMarketCommentFunc != nil {
		scenes.RestoreMarketCommentFunc(bot, comment)
	}
}

func deleteMarketComment(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	comment := update.DeleteMarketComment{}
	if err := json.Unmarshal(thisUpdate.Object, &comment); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketCommentDelete error is %s", err.Error()))
		return
	}

	if scenes.DeleteMarketCommentFunc != nil {
		scenes.DeleteMarketCommentFunc(bot, comment)
	}
}

func newMarketOrder(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	order := update.OrderMarket{}
	if err := json.Unmarshal(thisUpdate.Object, &order); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketOrderNew error is %s", err.Error()))
		return
	}

	if scenes.OrderMarketFunc != nil {
		scenes.OrderMarketFunc(bot, order)
	}
}

func editMarketOrder(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	order := update.OrderMarket{}
	if err := json.Unmarshal(thisUpdate.Object, &order); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("MarketOrderEdit error is %s", err.Error()))
		return
	}

	if scenes.EditOrderMarketFunc != nil {
		scenes.EditOrderMarketFunc(bot, order)
	}
}

func joinGroup(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	group := update.JoinGroup{}
	if err := json.Unmarshal(thisUpdate.Object, &group); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("GroupJoin error is %s", err.Error()))
		return
	}

	if scenes.JoinGroupFunc != nil {
		scenes.JoinGroupFunc(bot, group)
	}
}

func leaveGroup(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	group := update.LeaveGroup{}
	if err := json.Unmarshal(thisUpdate.Object, &group); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("GroupLeave error is %s", err.Error()))
		return
	}

	if scenes.LeaveGroupFunc != nil {
		scenes.LeaveGroupFunc(bot, group)
	}
}

func blockUser(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	user := update.BlockUser{}
	if err := json.Unmarshal(thisUpdate.Object, &user); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("UserBlock error is %s", err.Error()))
		return
	}

	if scenes.BlockUserFunc != nil {
		scenes.BlockUserFunc(bot, user)
	}
}

func unblockUser(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	user := update.UnblockUser{}
	if err := json.Unmarshal(thisUpdate.Object, &user); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("UserUnblock error is %s", err.Error()))
		return
	}

	if scenes.UnblockUserFunc != nil {
		scenes.UnblockUserFunc(bot, user)
	}
}

func newVote(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	vote := update.Vote{}
	if err := json.Unmarshal(thisUpdate.Object, &vote); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("PollVoteNew error is %s", err.Error()))
		return
	}

	if scenes.VoteFunc != nil {
		scenes.VoteFunc(bot, vote)
	}
}

func editOwners(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	owners := update.EditOwners{}
	if err := json.Unmarshal(thisUpdate.Object, &owners); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("GroupOfficersEdit error is %s", err.Error()))
		return
	}

	if scenes.EditOwnersFunc != nil {
		scenes.EditOwnersFunc(bot, owners)
	}
}

func controlGroup(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	group := update.ControlGroup{}
	if err := json.Unmarshal(thisUpdate.Object, &group); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("GroupChangeSettings error is %s", err.Error()))
		return
	}

	if scenes.ControlGroupFunc != nil {
		scenes.ControlGroupFunc(bot, group)
	}
}

func changeGroupPhoto(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	change := update.ChangeGroupPhoto{}
	if err := json.Unmarshal(thisUpdate.Object, &change); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("GroupChangePhoto error is %s", err.Error()))
		return
	}

	if scenes.ChangeGroupPhotoFunc != nil {
		scenes.ChangeGroupPhotoFunc(bot, change)
	}
}

func vkPayTransaction(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	transaction := update.VKPayTransaction{}
	if err := json.Unmarshal(thisUpdate.Object, &transaction); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("VKPayTransaction error is %s", err.Error()))
		return
	}

	if scenes.VKPayTransactionFunc != nil {
		scenes.VKPayTransactionFunc(bot, transaction)
	}
}

func appNotification(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	notification := update.AppNotification{}
	if err := json.Unmarshal(thisUpdate.Object, &notification); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("AppPayload error is %s", err.Error()))
		return
	}

	if scenes.AppNotificationFunc != nil {
		scenes.AppNotificationFunc(bot, notification)
	}
}

func donate(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	donate := update.ProlongedDonate{}
	if err := json.Unmarshal(thisUpdate.Object, &donate); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutSubscriptionCreate error is %s", err.Error()))
		return
	}

	if scenes.DonateFunc != nil {
		scenes.DonateFunc(bot, donate)
	}
}

func prolongedDonate(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	donate := update.ProlongedDonate{}
	if err := json.Unmarshal(thisUpdate.Object, &donate); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutSubscriptionProlonged error is %s", err.Error()))
		return
	}

	if scenes.ProlongedDonateFunc != nil {
		scenes.ProlongedDonateFunc(bot, donate)
	}
}

func expiredDonate(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	donate := update.ExpiredDonate{}
	if err := json.Unmarshal(thisUpdate.Object, &donate); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutSubscriptionExpired error is %s", err.Error()))
		return
	}

	if scenes.ExpiredDonateFunc != nil {
		scenes.ExpiredDonateFunc(bot, donate)
	}
}

func cancelDonate(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	donate := update.ExpiredDonate{}
	if err := json.Unmarshal(thisUpdate.Object, &donate); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutSubscriptionCancelled error is %s", err.Error()))
		return
	}

	if scenes.CancelDonateFunc != nil {
		scenes.CancelDonateFunc(bot, donate)
	}
}

func changeDonatePrice(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	donate := update.ChangeDonatePrice{}
	if err := json.Unmarshal(thisUpdate.Object, &donate); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutSubscriptionPriceChanged error is %s", err.Error()))
		return
	}

	if scenes.ChangeDonatePriceFunc != nil {
		scenes.ChangeDonatePriceFunc(bot, donate)
	}
}

func withdrawMoney(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	money := update.WithdrawMoney{}
	if err := json.Unmarshal(thisUpdate.Object, &money); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutMoneyWithdraw error is %s", err.Error()))
		return
	}

	if scenes.WithdrawMoneyFunc != nil {
		scenes.WithdrawMoneyFunc(bot, money)
	}
}

func withdrawMoneyError(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	money := update.WithdrawMoneyError{}
	if err := json.Unmarshal(thisUpdate.Object, &money); err != nil {
		bot.Bot.Logger.Error(fmt.Sprintf("DonutMoneyWithdrawError error is %s", err.Error()))
		return
	}

	if scenes.WithdrawMoneyErrorFunc != nil {
		scenes.WithdrawMoneyErrorFunc(bot, money)
	}
}

func Use(bot *api.Bot, thisUpdate update.Update, scenes Scenes) {
	if bot.Bot.Logger != nil {
		event, _ := json.Marshal(thisUpdate)
		bot.Bot.Logger.Info(fmt.Sprintf("Updates response is %s", string(event)))
	}

	handler := getScenesHandlers(thisUpdate.Type)
	if handler != nil {
		handler(bot, thisUpdate, scenes)
	}
}

func getScenesHandlers(typ string) Func {
	scenesHandlers := map[string]Func{
		"message_new":                      newMessage,
		"message_reply":                    replyMessage,
		"message_edit":                     editMessage,
		"message_allow":                    enableMessage,
		"message_deny":                     disableMessage,
		"message_typing_state":             typingMessage,
		"message_event":                    eventMessage,
		"photo_new":                        newPhoto,
		"photo_comment_new":                newPhotoComment,
		"photo_comment_edit":               editPhotoComment,
		"photo_comment_restore":            restorePhotoComment,
		"photo_comment_delete":             deletePhotoComment,
		"audio_new":                        newAudio,
		"video_new":                        newVideo,
		"video_comment_new":                newVideoComment,
		"video_comment_edit":               editVideoComment,
		"video_comment_restore":            restoreVideoComment,
		"video_comment_delete":             deleteVideoComment,
		"wall_post_new":                    newWallPost,
		"wall_repost":                      repostWall,
		"wall_reply_new":                   newWallReply,
		"wall_reply_edit":                  editWallReply,
		"wall_reply_restore":               restoreWallReply,
		"wall_reply_delete":                deleteWallReply,
		"like_add":                         like,
		"like_remove":                      unlike,
		"board_post_new":                   newBoardPost,
		"board_post_edit":                  editBoardPost,
		"board_post_restore":               restoreBoardPost,
		"board_post_delete":                deleteBoardPost,
		"market_comment":                   marketComment,
		"market_comment_edit":              editMarketComment,
		"market_comment_restore":           restoreMarketComment,
		"market_comment_delete":            deleteMarketComment,
		"market_order_new":                 newMarketOrder,
		"market_order_edit":                editMarketOrder,
		"group_join":                       joinGroup,
		"group_leave":                      leaveGroup,
		"user_block":                       blockUser,
		"user_unblock":                     unblockUser,
		"poll_vote_new":                    newVote,
		"group_officers_edit":              editOwners,
		"group_change_settings":            controlGroup,
		"group_change_photo":               changeGroupPhoto,
		"vkpay_transaction":                vkPayTransaction,
		"app_payload":                      appNotification,
		"donut_subscription_create":        donate,
		"donut_subscription_prolonged":     prolongedDonate,
		"donut_subscription_expired":       expiredDonate,
		"donut_subscription_cancelled":     cancelDonate,
		"donut_subscription_price_changed": changeDonatePrice,
		"donut_money_withdraw":             withdrawMoney,
		"donut_money_withdraw_error":       withdrawMoneyError,
	}

	return scenesHandlers[typ]
}
