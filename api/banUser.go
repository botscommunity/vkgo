package api

import (
	"fmt"
	"net/url"
	"time"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) BanUser(properties ...any) (ok bool, err *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.BanUser:
			setBanUserStructProperty(&query, property)
		case int:
			setBanUserIntProperty(&query, property)
		case string:
			query.Set("comment", property)
		case bool:
			query.Set("comment_visible", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	reply := responses.Ban{}
	err = responses.NewInternalError(bot.Call("groups.ban", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func setBanUserStructProperty(query *url.Values, properties types.BanUser) {
	if properties.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(properties.GroupID))
	}

	if properties.UserID != 0 {
		query.Set("owner_id", fmt.Sprint(properties.UserID))
	}

	if properties.Time != 0 {
		query.Set("end_date", fmt.Sprint(time.Now().Unix()+int64(properties.Time)))
	}

	if properties.Reason != 0 {
		query.Set("reason", fmt.Sprint(properties.Reason))
	}

	if properties.Comment.Text != "" {
		query.Set("comment", properties.Comment.Text)
	}

	if !query.Has("comment_visible") && properties.Comment.Show {
		query.Set("comment_visible", fmt.Sprint(conv.BooleanToInteger(properties.Comment.Show)))
	}
}

func setBanUserIntProperty(query *url.Values, value int) {
	switch {
	case !query.Has("group_id"):
		query.Set("group_id", fmt.Sprint(value))
	case !query.Has("owner_id"):
		query.Set("owner_id", fmt.Sprint(value))
	default:
		query.Set("end_date", fmt.Sprint(time.Now().Unix()+int64(value)))
	}
}

func (bot *Bot) UnbanUser(properties ...any) (ok bool, err *responses.Error) {
	query, groupID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.UnbanUser:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))

				groupID = true
			}

			if property.UserID != 0 {
				query.Set("owner_id", fmt.Sprint(property.UserID))
			}
		case int:
			if !groupID {
				query.Set("group_id", fmt.Sprint(property))

				groupID = true
			} else {
				query.Set("owner_id", fmt.Sprint(property))
			}
		}
	}

	reply := responses.Ban{}
	err = responses.NewInternalError(bot.Call("groups.unban", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func (bot *Bot) GetBannedUsers(properties ...any) (count int, bans []responses.BannedUser, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.BannedUsers:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))

				installed["groupID"] = true
			}

			if property.UserID != 0 {
				query.Set("owner_id", fmt.Sprint(property.UserID))

				installed["userID"] = true
			}

			if property.Offset != 0 {
				query.Set("offset", fmt.Sprint(property.Offset))

				installed["offset"] = true
			}

			if property.Count != 0 {
				query.Set("count", fmt.Sprint(property.Count))
			}

			if len(property.Fields) != 0 {
				query.Set("fields", conv.ForParams(property.Fields))
			}
		case int:
			switch {
			case !installed["groupID"]:
				query.Set("group_id", fmt.Sprint(property))

				installed["groupID"] = true
			case !installed["userID"]:
				query.Set("owner_id", fmt.Sprint(property))

				installed["userID"] = true
			case !installed["offset"]:
				query.Set("offset", fmt.Sprint(property))

				installed["offset"] = true
			default:
				query.Set("count", fmt.Sprint(property))
			}
		case []string:
			query.Set("fields", conv.ForParams(property))
		}
	}

	reply := responses.BannedUsersReply{}
	err = responses.NewInternalError(bot.Call("groups.getBanned", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Count, reply.Response.Bans, err
}
