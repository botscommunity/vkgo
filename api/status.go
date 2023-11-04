package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetStatus(properties ...any) (text string, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.Status:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))
			}

			if property.UserID != 0 {
				query.Set("user_id", fmt.Sprint(property.UserID))

				installed["userID"] = true
			}
		case int:
			if !installed["userID"] {
				query.Set("user_id", fmt.Sprint(property))

				installed["userID"] = true
			} else {
				query.Set("group_id", fmt.Sprint(property))
			}
		}
	}

	reply := responses.StatusReply{}
	err = responses.NewInternalError(bot.Call("status.get", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Text, err
}

func (bot *Bot) SetStatus(properties ...any) (ok bool, err *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.SetStatus:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))
			}

			if property.Text != "" {
				query.Set("text", property.Text)
			}
		case int:
			query.Set("group_id", fmt.Sprint(property))
		case string:
			query.Set("text", property)
		}
	}

	reply := responses.SetStatus{}
	err = responses.NewInternalError(bot.Call("status.set", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}
