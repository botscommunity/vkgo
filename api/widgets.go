package api

import (
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) UpdateWidgets(properties ...any) (ok bool, err *responses.Error) {
	query, typ := make(url.Values), false

	for _, property := range properties {
		if property, ok := property.(string); ok {
			if !typ {
				query.Set("type", property)

				typ = true
			} else {
				query.Set("code", property)
			}
		}
	}

	reply := responses.UpdateWidgets{}
	err = responses.NewInternalError(bot.Call("appWidgets.update", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Response == 1 {
		ok = true
	}

	return
}
