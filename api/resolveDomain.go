package api

import (
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) ResolveDomain(properties ...any) (typ string, objectID int, err *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.ResolveDomain:
			if property.Domain != "" {
				query.Set("screen_name", property.Domain)
			}
		case string:
			query.Set("screen_name", property)
		}
	}

	reply := responses.ResolveDomainReply{}
	err = responses.NewInternalError(bot.Call("utils.resolveScreenName", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Type, reply.Response.ObjectID, err
}
