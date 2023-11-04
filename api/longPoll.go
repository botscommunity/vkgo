package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetUserLongPollServer(properties ...any) responses.LongPollServer {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.UserLongPollServer:
			if property.Version != 0 {
				query.Set("lp_version", fmt.Sprint(property.Version))

				installed["version"] = true
			}

			if property.GroupID > 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))

				installed["groupID"] = true
			}

			if property.NeedPTS {
				query.Set("need_pts", fmt.Sprint(conv.BooleanToInteger(property.NeedPTS)))
			}
		case int:
			if !installed["version"] {
				query.Set("lp_version", fmt.Sprint(property))

				installed["version"] = true
			} else if !installed["groupID"] {
				query.Set("group_id", fmt.Sprint(property))

				installed["groupID"] = true
			}
		case bool:
			query.Set("need_pts", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	var (
		reply = responses.LongPollServerReply{}
		err   = responses.NewInternalError(bot.Call("messages.getLongPollServer", query.Encode(), &reply))
	)

	if err != nil {
		reply.Response.Error = err
		return reply.Response
	}

	if reply.Error != nil {
		reply.Response.Error = reply.Error
	}

	return reply.Response
}

func (bot *Bot) GetGroupLongPollServer(properties ...any) responses.LongPollServer {
	query, installed := url.Values{}, make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.GroupLongPollServer:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))

				installed["groupID"] = true
			}
		case int:
			query.Set("group_id", fmt.Sprint(property))

			installed["groupID"] = true
		}

		if bot.IsGroup {
			query.Set("group_id", fmt.Sprint(bot.ID))

			installed["groupID"] = true
		}
	}

	if !installed["groupID"] {
		query.Set("group_id", fmt.Sprint(bot.ID))
	}

	var (
		reply = responses.LongPollServerReply{}
		err   = responses.NewInternalError(bot.Call("groups.getLongPollServer", query.Encode(), &reply))
	)

	if err != nil {
		reply.Response.Error = err
		return reply.Response
	}

	if reply.Error != nil {
		reply.Response.Error = reply.Error
	}

	return reply.Response
}
