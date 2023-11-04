package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) KickUser(properties ...any) (ok bool, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.KickUser:
			if property.ChatID != 0 {
				query.Set("chat_id", fmt.Sprint(property.ChatID))

				installed["chat_id"] = true
			}

			if property.MemberID != 0 {
				query.Set("member_id", fmt.Sprint(property.MemberID))

				installed["member_id"] = true
			}

			if property.UserID != 0 {
				query.Set("user_id", fmt.Sprint(property.UserID))
			}
		case int:
			if !installed["chat_id"] {
				query.Set("chat_id", fmt.Sprint(property))

				installed["chat_id"] = true
			} else {
				query.Set("member_id", fmt.Sprint(property))
			}
		}
	}

	reply := responses.KickUser{}
	err = responses.NewInternalError(bot.Call("messages.removeChatUser", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}
