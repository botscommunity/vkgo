package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"

	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) SendAnswerEvent(properties ...any) (ok bool, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.SendEventMessage:
			if property.EventID != "" {
				query.Set("event_id", property.EventID)

				installed["eventID"] = true
			}

			if property.ChatID != 0 {
				query.Set("peer_id", fmt.Sprint(property.ChatID))

				installed["chatID"] = true
			}

			if property.UserID != 0 {
				query.Set("user_id", fmt.Sprint(property.UserID))
			}

			if property.Action != "" {
				query.Set("event_data", property.Action)
			}
		case types.EventAction:
			query.Set("event_data", property.JSON())
		case int:
			if !installed["chatID"] {
				query.Set("peer_id", fmt.Sprint(property))

				installed["chatID"] = true
			} else {
				query.Set("user_id", fmt.Sprint(property))
			}
		case string:
			if !installed["eventID"] {
				query.Set("event_id", property)

				installed["eventID"] = true
			} else {
				query.Set("event_data", property)
			}
		}
	}

	reply := responses.SendEventMessage{}
	err = responses.NewInternalError(bot.Call("messages.sendMessageEventAnswer", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}
