package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetGroup(properties ...any) (group responses.Group) {
	groups, err := bot.Groups(properties...)

	if len(groups) > 0 {
		group = groups[0]
	}

	if err != nil {
		group.Error = err
	}

	return
}

func (bot *Bot) Groups(properties ...any) (groups []responses.Group, err *responses.Error) {
	query, groupID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.Group:
			if property.GroupID != 0 {
				query.Set("group_ids", fmt.Sprint(property.GroupID))

				groupID = true
			} else if len(property.GroupsID) != 0 {
				query.Set("group_ids", conv.ForParams(property.GroupsID))

				groupID = true
			}

			if len(property.Fields) != 0 {
				query.Set("fields", conv.ForParams(property.Fields))
			}
		case int:
			query.Set("group_ids", fmt.Sprint(property))

			groupID = true
		case []int:
			query.Set("group_ids", conv.ForParams(property))

			groupID = true
		case []string:
			query.Set("fields", conv.ForParams(property))
		}
	}

	if !groupID && bot.IsGroup {
		query.Set("group_ids", fmt.Sprint(bot.ID))
	}

	reply := responses.GroupsReply{}
	err = responses.NewInternalError(bot.Call("groups.getById", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Groups, err
}
