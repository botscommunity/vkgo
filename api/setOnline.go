package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) SetOnline(properties ...any) (ok bool, err *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.SetOnline:
			if property.Calls {
				query.Set("voip", fmt.Sprint(conv.BooleanToInteger(property.Calls)))
			}
		case int:
			query.Set("voip", fmt.Sprint(property))
		case bool:
			query.Set("voip", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	reply := responses.OnlineStatus{}
	err = responses.NewInternalError(bot.Call("account.setOnline", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return
}

func (bot *Bot) SetOffline() (ok bool, err *responses.Error) {
	reply := responses.OnlineStatus{}
	err = responses.NewInternalError(bot.Call("account.setOffline", "", &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return
}
