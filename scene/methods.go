package scene

// Messages

func NewMessage(msgFunc MessageFunc) *Scenes {
	return New().Message(msgFunc)
}

func (scenes *Scenes) NewMessage(msgFunc MessageFunc) *Scenes {
	return scenes.Message(msgFunc)
}

func Message(msgFunc MessageFunc) *Scenes {
	return New().Message(msgFunc)
}

func (scenes *Scenes) Message(msgFunc MessageFunc) *Scenes {
	scenes.MessageFunc = msgFunc
	return scenes
}

func Callback(cbFunc CallbackFunc) *Scenes {
	return New().Callback(cbFunc)
}

func (scenes *Scenes) Callback(cbFunc CallbackFunc) *Scenes {
	scenes.CallbackFunc = cbFunc
	return scenes
}

func ReplyMessage(replMsgFunc ReplyMessageFunc) *Scenes {
	return New().ReplyMessage(replMsgFunc)
}

func (scenes *Scenes) ReplyMessage(replMsgFunc ReplyMessageFunc) *Scenes {
	scenes.ReplyMessageFunc = replMsgFunc
	return scenes
}

func EditMessage(edMsgFunc EditMessageFunc) *Scenes {
	return New().EditMessage(edMsgFunc)
}

func (scenes *Scenes) EditMessage(edMsgFunc EditMessageFunc) *Scenes {
	scenes.EditMessageFunc = edMsgFunc
	return scenes
}

func EnableMessage(enbMsgFunc EnableMessageFunc) *Scenes {
	return New().EnableMessage(enbMsgFunc)
}

func (scenes *Scenes) EnableMessage(enbMsgFunc EnableMessageFunc) *Scenes {
	scenes.EnableMessageFunc = enbMsgFunc
	return scenes
}

func DisableMessage(disMsgFunc DisableMessageFunc) *Scenes {
	return New().DisableMessage(disMsgFunc)
}

func (scenes *Scenes) DisableMessage(disMsgFunc DisableMessageFunc) *Scenes {
	scenes.DisableMessageFunc = disMsgFunc
	return scenes
}

func TypingMessage(typFunc TypingMessageFunc) *Scenes {
	return New().TypingMessage(typFunc)
}

func (scenes *Scenes) TypingMessage(typFunc TypingMessageFunc) *Scenes {
	scenes.TypingMessageFunc = typFunc
	return scenes
}

// Photos

func Photo(photoFunc PhotoFunc) *Scenes {
	return New().Photo(photoFunc)
}

func (scenes *Scenes) Photo(photoFunc PhotoFunc) *Scenes {
	scenes.PhotoFunc = photoFunc
	return scenes
}

func PhotoComment(photoComFunc PhotoCommentFunc) *Scenes {
	return New().PhotoComment(photoComFunc)
}

func (scenes *Scenes) PhotoComment(photoComFunc PhotoCommentFunc) *Scenes {
	scenes.PhotoCommentFunc = photoComFunc
	return scenes
}

func EditPhotoComment(edPhotoFunc EditPhotoCommentFunc) *Scenes {
	return New().EditPhotoComment(edPhotoFunc)
}

func (scenes *Scenes) EditPhotoComment(edPhotoFunc EditPhotoCommentFunc) *Scenes {
	scenes.EditPhotoCommentFunc = edPhotoFunc
	return scenes
}

func RestorePhotoComment(restPhotoComFunc RestorePhotoCommentFunc) *Scenes {
	return New().RestorePhotoComment(restPhotoComFunc)
}

func (scenes *Scenes) RestorePhotoComment(restPhotoComFunc RestorePhotoCommentFunc) *Scenes {
	scenes.RestorePhotoCommentFunc = restPhotoComFunc
	return scenes
}

func DeletePhotoComment(delPhotoComFunc DeletePhotoCommentFunc) *Scenes {
	return New().DeletePhotoComment(delPhotoComFunc)
}

func (scenes *Scenes) DeletePhotoComment(delPhotoComFunc DeletePhotoCommentFunc) *Scenes {
	scenes.DeletePhotoCommentFunc = delPhotoComFunc
	return scenes
}

// Sounds

func Audio(audFunc NewAudioFunc) *Scenes {
	return New().Audio(audFunc)
}

func (scenes *Scenes) Audio(audFunc NewAudioFunc) *Scenes {
	scenes.NewAudioFunc = audFunc
	return scenes
}

// Videos

func Video(vidFunc NewVideoFunc) *Scenes {
	return New().Video(vidFunc)
}

func (scenes *Scenes) Video(vidFunc NewVideoFunc) *Scenes {
	scenes.NewVideoFunc = vidFunc
	return scenes
}

func VideoComment(vidComFunc VideoCommentFunc) *Scenes {
	return New().VideoComment(vidComFunc)
}

func (scenes *Scenes) VideoComment(vidComFunc VideoCommentFunc) *Scenes {
	scenes.VideoCommentFunc = vidComFunc
	return scenes
}

func EditVideoComment(edVidComFunc EditVideoCommentFunc) *Scenes {
	return New().EditVideoComment(edVidComFunc)
}

func (scenes *Scenes) EditVideoComment(edVidComFunc EditVideoCommentFunc) *Scenes {
	scenes.EditVideoCommentFunc = edVidComFunc
	return scenes
}

func RestoreVideoComment(restVidComFunc RestoreVideoCommentFunc) *Scenes {
	return New().RestoreVideoComment(restVidComFunc)
}

func (scenes *Scenes) RestoreVideoComment(restVidComFunc RestoreVideoCommentFunc) *Scenes {
	scenes.RestoreVideoCommentFunc = restVidComFunc
	return scenes
}

func DeleteVideoComment(delVidComFunc DeleteVideoCommentFunc) *Scenes {
	return New().DeleteVideoComment(delVidComFunc)
}

func (scenes *Scenes) DeleteVideoComment(delVidComFunc DeleteVideoCommentFunc) *Scenes {
	scenes.DeleteVideoCommentFunc = delVidComFunc
	return scenes
}

// Posts

func Post(postFunc PostFunc) *Scenes {
	return New().Post(postFunc)
}

func (scenes *Scenes) Post(postFunc PostFunc) *Scenes {
	scenes.PostFunc = postFunc
	return scenes
}

func Repost(postFunc RepostFunc) *Scenes {
	return New().Repost(postFunc)
}

func (scenes *Scenes) Repost(repost RepostFunc) *Scenes {
	scenes.RepostFunc = repost
	return scenes
}

func PostComment(postComFunc PostCommentFunc) *Scenes {
	return New().PostComment(postComFunc)
}

func (scenes *Scenes) PostComment(postComFunc PostCommentFunc) *Scenes {
	scenes.PostCommentFunc = postComFunc
	return scenes
}

func EditPostComment(edPostComFunc EditPostCommentFunc) *Scenes {
	return New().EditPostComment(edPostComFunc)
}

func (scenes *Scenes) EditPostComment(editPostComment EditPostCommentFunc) *Scenes {
	scenes.EditPostCommentFunc = editPostComment
	return scenes
}

func RestorePostComment(restPostFunc RestorePostCommentFunc) *Scenes {
	return New().RestorePostComment(restPostFunc)
}

func (scenes *Scenes) RestorePostComment(restPostFunc RestorePostCommentFunc) *Scenes {
	scenes.RestorePostCommentFunc = restPostFunc
	return scenes
}

func DeletePostComment(delPostComFunc DeletePostCommentFunc) *Scenes {
	return New().DeletePostComment(delPostComFunc)
}

func (scenes *Scenes) DeletePostComment(delPostComFunc DeletePostCommentFunc) *Scenes {
	scenes.DeletePostCommentFunc = delPostComFunc
	return scenes
}

// Likes

func Like(likeFunc LikeFunc) *Scenes {
	return New().Like(likeFunc)
}

func (scenes *Scenes) Like(likeFunc LikeFunc) *Scenes {
	scenes.LikeFunc = likeFunc
	return scenes
}

func Unlike(unlikeFunc UnlikeFunc) *Scenes {
	return New().Unlike(unlikeFunc)
}

func (scenes *Scenes) Unlike(unlikeFunc UnlikeFunc) *Scenes {
	scenes.UnlikeFunc = unlikeFunc
	return scenes
}

// Board Comments

func BoardComment(boardComFunc BoardCommentFunc) *Scenes {
	return New().BoardComment(boardComFunc)
}

func (scenes *Scenes) BoardComment(boardComFunc BoardCommentFunc) *Scenes {
	scenes.BoardCommentFunc = boardComFunc
	return scenes
}

func EditBoardComment(edBoardComFunc EditBoardCommentFunc) *Scenes {
	return New().EditBoardComment(edBoardComFunc)
}

func (scenes *Scenes) EditBoardComment(edBoardComFunc EditBoardCommentFunc) *Scenes {
	scenes.EditBoardCommentFunc = edBoardComFunc
	return scenes
}

func RestoreBoardComment(restBoardComFunc RestoreBoardCommentFunc) *Scenes {
	return New().RestoreBoardComment(restBoardComFunc)
}

func (scenes *Scenes) RestoreBoardComment(restBoardComFunc RestoreBoardCommentFunc) *Scenes {
	scenes.RestoreBoardCommentFunc = restBoardComFunc
	return scenes
}

func DeleteBoardComment(delBoardComFunc DeleteBoardCommentFunc) *Scenes {
	return New().DeleteBoardComment(delBoardComFunc)
}

func (scenes *Scenes) DeleteBoardComment(delBoardComFunc DeleteBoardCommentFunc) *Scenes {
	scenes.DeleteBoardCommentFunc = delBoardComFunc
	return scenes
}

// Market Comments

func MarketComment(marketComFunc MarketCommentFunc) *Scenes {
	return New().MarketComment(marketComFunc)
}

func (scenes *Scenes) MarketComment(marketComFunc MarketCommentFunc) *Scenes {
	scenes.MarketCommentFunc = marketComFunc
	return scenes
}

func EditMarketComment(edMarketComFunc EditMarketCommentFunc) *Scenes {
	return New().EditMarketComment(edMarketComFunc)
}

func (scenes *Scenes) EditMarketComment(edMarketComFunc EditMarketCommentFunc) *Scenes {
	scenes.EditMarketCommentFunc = edMarketComFunc
	return scenes
}

func RestoreMarketComment(restMarketComFunc RestoreMarketCommentFunc) *Scenes {
	return New().RestoreMarketComment(restMarketComFunc)
}

func (scenes *Scenes) RestoreMarketComment(restMarketComFunc RestoreMarketCommentFunc) *Scenes {
	scenes.RestoreMarketCommentFunc = restMarketComFunc
	return scenes
}

func DeleteMarketComment(delMarketComFunc DeleteMarketCommentFunc) *Scenes {
	return New().DeleteMarketComment(delMarketComFunc)
}

func (scenes *Scenes) DeleteMarketComment(delMarketComFunc DeleteMarketCommentFunc) *Scenes {
	scenes.DeleteMarketCommentFunc = delMarketComFunc
	return scenes
}

// Orders

func OrderMarket(orderMarketFunc OrderMarketFunc) *Scenes {
	return New().OrderMarket(orderMarketFunc)
}

func (scenes *Scenes) OrderMarket(orderMarketFunc OrderMarketFunc) *Scenes {
	scenes.OrderMarketFunc = orderMarketFunc
	return scenes
}

func EditOrderMarket(edOrderMarketFunc EditOrderMarketFunc) *Scenes {
	return New().EditOrderMarket(edOrderMarketFunc)
}

func (scenes *Scenes) EditOrderMarket(edOrderMarketFunc EditOrderMarketFunc) *Scenes {
	scenes.EditOrderMarketFunc = edOrderMarketFunc
	return scenes
}

func EditOrder(edOrderMarketFunc EditOrderMarketFunc) *Scenes {
	return New().EditOrder(edOrderMarketFunc)
}

func (scenes *Scenes) EditOrder(edOrderMarketFunc EditOrderMarketFunc) *Scenes {
	scenes.EditOrderMarketFunc = edOrderMarketFunc
	return scenes
}

// Community

func JoinGroup(jGroupFunc JoinGroupFunc) *Scenes {
	return New().JoinGroup(jGroupFunc)
}

func (scenes *Scenes) JoinGroup(jGroupFunc JoinGroupFunc) *Scenes {
	scenes.JoinGroupFunc = jGroupFunc
	return scenes
}

func LeaveGroup(lGroupFunc LeaveGroupFunc) *Scenes {
	return New().LeaveGroup(lGroupFunc)
}

func (scenes *Scenes) LeaveGroup(lGroupFunc LeaveGroupFunc) *Scenes {
	scenes.LeaveGroupFunc = lGroupFunc
	return scenes
}

func BlockUser(blockUserFunc BlockUserFunc) *Scenes {
	return New().BlockUser(blockUserFunc)
}

func (scenes *Scenes) BlockUser(blockUserFunc BlockUserFunc) *Scenes {
	scenes.BlockUserFunc = blockUserFunc
	return scenes
}

func UnblockUser(unblockUserFunc UnblockUserFunc) *Scenes {
	return New().UnblockUser(unblockUserFunc)
}

func (scenes *Scenes) UnblockUser(unblockUserFunc UnblockUserFunc) *Scenes {
	scenes.UnblockUserFunc = unblockUserFunc
	return scenes
}

// Others

func Vote(voteFunc VoteFunc) *Scenes {
	return New().Vote(voteFunc)
}

func (scenes *Scenes) Vote(voteFunc VoteFunc) *Scenes {
	scenes.VoteFunc = voteFunc
	return scenes
}

func EditOwners(edPwnFunc EditOwnersFunc) *Scenes {
	return New().EditOwners(edPwnFunc)
}

func (scenes *Scenes) EditOwners(edPwnFunc EditOwnersFunc) *Scenes {
	scenes.EditOwnersFunc = edPwnFunc
	return scenes
}

func ControlGroup(ctrlGroupFunc ControlGroupFunc) *Scenes {
	return New().ControlGroup(ctrlGroupFunc)
}

func (scenes *Scenes) ControlGroup(ctrlGroupFunc ControlGroupFunc) *Scenes {
	scenes.ControlGroupFunc = ctrlGroupFunc
	return scenes
}

func ChangeGroupPhoto(chGroupPhotoFunc ChangeGroupPhotoFunc) *Scenes {
	return New().ChangeGroupPhoto(chGroupPhotoFunc)
}

func (scenes *Scenes) ChangeGroupPhoto(chGroupPhotoFunc ChangeGroupPhotoFunc) *Scenes {
	scenes.ChangeGroupPhotoFunc = chGroupPhotoFunc
	return scenes
}

func VKPayTransaction(vkPayTransFunc VKPayTransactionFunc) *Scenes {
	return New().VKPayTransaction(vkPayTransFunc)
}

func (scenes *Scenes) VKPayTransaction(vkPayTransFunc VKPayTransactionFunc) *Scenes {
	scenes.VKPayTransactionFunc = vkPayTransFunc
	return scenes
}

func AppNotification(appNotifyFunc AppNotificationFunc) *Scenes {
	return New().ApplicationNotification(appNotifyFunc)
}

func ApplicationNotification(appNotifyFunc AppNotificationFunc) *Scenes {
	return New().ApplicationNotification(appNotifyFunc)
}

func (scenes *Scenes) ApplicationNotification(appNotifyFunc AppNotificationFunc) *Scenes {
	scenes.AppNotificationFunc = appNotifyFunc
	return scenes
}

func Donate(donFunc DonateFunc) *Scenes {
	return New().Donate(donFunc)
}

func (scenes *Scenes) Donate(donFunc DonateFunc) *Scenes {
	scenes.DonateFunc = donFunc
	return scenes
}

func ProlongedDonate(proDonFunc ProlongedDonateFunc) *Scenes {
	return New().ProlongedDonate(proDonFunc)
}

func (scenes *Scenes) ProlongedDonate(proDonFunc ProlongedDonateFunc) *Scenes {
	scenes.ProlongedDonateFunc = proDonFunc
	return scenes
}

func ExpiredDonate(expDonFunc ExpiredDonateFunc) *Scenes {
	return New().ExpiredDonate(expDonFunc)
}

func (scenes *Scenes) ExpiredDonate(expDonFunc ExpiredDonateFunc) *Scenes {
	scenes.ExpiredDonateFunc = expDonFunc
	return scenes
}

func CancelDonate(cancDonFunc CancelDonateFunc) *Scenes {
	return New().CancelDonate(cancDonFunc)
}

func (scenes *Scenes) CancelDonate(cancDonFunc CancelDonateFunc) *Scenes {
	scenes.CancelDonateFunc = cancDonFunc
	return scenes
}

func ChangeDonatePrice(chDonPriceFunc ChangeDonatePriceFunc) *Scenes {
	return New().ChangeDonatePrice(chDonPriceFunc)
}

func (scenes *Scenes) ChangeDonatePrice(chDonPriceFunc ChangeDonatePriceFunc) *Scenes {
	scenes.ChangeDonatePriceFunc = chDonPriceFunc
	return scenes
}

func WithdrawMoney(withdrawMoneyFunc WithdrawMoneyFunc) *Scenes {
	return New().WithdrawMoney(withdrawMoneyFunc)
}

func (scenes *Scenes) WithdrawMoney(withdrawMoneyFunc WithdrawMoneyFunc) *Scenes {
	scenes.WithdrawMoneyFunc = withdrawMoneyFunc
	return scenes
}

func WithdrawMoneyError(withdrawMoneyErrFunc WithdrawMoneyErrorFunc) *Scenes {
	return New().WithdrawMoneyError(withdrawMoneyErrFunc)
}

func (scenes *Scenes) WithdrawMoneyError(withdrawMoneyErrFunc WithdrawMoneyErrorFunc) *Scenes {
	scenes.WithdrawMoneyErrorFunc = withdrawMoneyErrFunc
	return scenes
}
