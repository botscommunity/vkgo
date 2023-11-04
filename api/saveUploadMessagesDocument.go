package api

import (
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) SaveUploadMessagesDocument(file string) (typ string, document responses.SaveUploadMessageDocument, err *responses.Error) {
	query := make(url.Values)
	query.Set("file", file)

	reply := responses.SaveUploadMessagesDocumentReply{}
	err = responses.NewInternalError(bot.Call("docs.save", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Type, reply.Response.Document, err
}
