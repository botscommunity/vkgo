package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) SaveUploadMessagesPhoto(server int, photo, hash string) (photos []responses.SaveUploadMessagesPhoto, err *responses.Error) {
	query := make(url.Values)
	query.Set("server", fmt.Sprint(server))
	query.Set("photo", photo)
	query.Set("hash", hash)

	reply := responses.SaveUploadMessagesPhotoReply{}
	err = responses.NewInternalError(bot.Call("photos.saveMessagesPhoto", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Photos, err
}
