package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) SetGroupOnline(properties ...any) (ok bool, err *responses.Error) {
	query, groupID := make(url.Values), false

	for _, property := range properties {
		if property, ok := property.(int); ok {
			query.Set("group_id", fmt.Sprint(property))

			groupID = true
		}
	}

	if !groupID {
		query.Set("group_id", fmt.Sprint(bot.ID))
	}

	reply := responses.OnlineStatus{}
	err = responses.NewInternalError(bot.Call("groups.enableOnline", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return
}

func (bot *Bot) SetGroupOffline(properties ...any) (ok bool, err *responses.Error) {
	query, groupID := make(url.Values), false

	for _, property := range properties {
		if property, ok := property.(int); ok {
			query.Set("group_id", fmt.Sprint(property))

			groupID = true
		}
	}

	if !groupID {
		query.Set("group_id", fmt.Sprint(bot.ID))
	}

	reply := responses.OnlineStatus{}
	err = responses.NewInternalError(bot.Call("groups.disableOnline", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return
}
