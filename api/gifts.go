package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetGifts(properties ...any) (count int, gifts []responses.Gift, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.Gifts:
			if property.UserID != 0 {
				query.Set("user_id", fmt.Sprintf("%d", property.UserID))

				installed["userID"] = true
			}

			if property.Offset != 0 {
				query.Set("offset", fmt.Sprintf("%d", property.Offset))

				installed["offset"] = true
			}

			if property.Count != 0 {
				query.Set("count", fmt.Sprintf("%d", property.Count))
			}
		case int:
			switch {
			case !installed["userID"]:
				query.Set("user_id", fmt.Sprintf("%d", property))

				installed["userID"] = true
			case !installed["offset"]:
				query.Set("offset", fmt.Sprintf("%d", property))

				installed["offset"] = true
			default:
				query.Set("count", fmt.Sprintf("%d", property))
			}
		}
	}

	reply := responses.GiftsReply{}
	err = responses.NewInternalError(bot.Call("gifts.get", query.Encode(), &gifts))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Count, reply.Response.Gifts, err
}
