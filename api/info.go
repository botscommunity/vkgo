package api

import (
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetInfo(properties ...any) responses.Info {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.Info:
			if len(property.Fields) > 0 {
				query.Set("fields", conv.ForParams(property.Fields))
			}
		case []string:
			query.Set("fields", conv.ForParams(property))
		}
	}

	var (
		reply = responses.InfoReply{}
		err   = responses.NewInternalError(bot.Call("account.getInfo", query.Encode(), &reply))
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
