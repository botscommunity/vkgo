package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) Ban(objectID int) (ok bool, err *responses.Error) {
	reply := responses.Ban{}
	if err = responses.NewInternalError(bot.Call("account.ban", fmt.Sprintf("owner_id=%d", objectID), &reply)); reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func (bot *Bot) Unban(objectID int) (ok bool, err *responses.Error) {
	reply := responses.Ban{}
	if err = responses.NewInternalError(bot.Call("account.unban", fmt.Sprintf("owner_id=%d", objectID), &reply)); reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func (bot *Bot) GetBanned(properties ...any) (count int, bans []int, users []responses.User, err *responses.Error) {
	query, offset := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.Banned:
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

	reply := responses.BannedReply{}
	err = responses.NewInternalError(bot.Call("account.getBanned", "", &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Count, reply.Response.Bans, reply.Response.Users, err
}
