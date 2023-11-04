package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetUploadMessagesPhotoServer(property any) responses.UploadMessagesPhotoServer {
	query := make(url.Values)

	switch property := property.(type) {
	case types.UploadMessagesPhotoServer:
		if property.ChatID != 0 {
			query.Set("peer_id", fmt.Sprint(property.ChatID))
		}
	case int:
		query.Set("peer_id", fmt.Sprint(property))
	}

	var (
		reply = responses.UploadMessagesPhotoServerReply{}
		err   = responses.NewInternalError(bot.Call("photos.getMessagesUploadServer", query.Encode(), &reply))
	)

	if err != nil {
		reply.Response.Error = err
		return reply.Response
	}

	if reply.Error != nil {
		reply.Response.Error = reply.Error
	}

	return reply.Response
}
