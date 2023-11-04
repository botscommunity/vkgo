package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetOffers(properties ...any) (count int, offers []int, err *responses.Error) {
	query, offset := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.Offers:
			if property.Offset != 0 {
				query.Set("offset", fmt.Sprint(property.Offset))

				offset = true
			}

			if property.Count != 0 {
				query.Set("count", fmt.Sprint(property.Count))
			}
		case int:
			if !offset {
				query.Set("offset", fmt.Sprint(property))

				offset = true
			} else {
				query.Set("count", fmt.Sprint(property))
			}
		}
	}

	reply := responses.OffersReply{}
	err = responses.NewInternalError(bot.Call("account.getActiveOffers", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Count, reply.Response.Offers, err
}
