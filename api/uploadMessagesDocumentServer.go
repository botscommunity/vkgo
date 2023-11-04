package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetUploadMessagesDocumentServer(properties ...any) responses.UploadMessagesDocumentServer {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.UploadMessagesDocumentServer:
			if property.ChatID != 0 {
				query.Set("peer_id", fmt.Sprint(property.ChatID))
			}

			if property.Type != "" {
				query.Set("type", property.Type)
			}
		case int:
			query.Set("peer_id", fmt.Sprint(property))
		case string:
			query.Set("type", property)
		}
	}

	var (
		reply = responses.UploadMessagesDocumentServerReply{}
		err   = responses.NewInternalError(bot.Call("docs.getMessagesUploadServer", query.Encode(), &reply))
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
