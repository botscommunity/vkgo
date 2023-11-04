package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) GetHistoryMessages(properties ...any) responses.HistoryMessages {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.HistoryMessages:
			setHistoryMessagesStructProperty(&query, property, installed)
		case int:
			setHistoryMessagesIntProperty(&query, property, installed)
		case []string:
			query.Set("fields", conv.ForParams(property))
		case bool:
			setHistoryMessagesBoolProperty(&query, property, installed)
		}
	}

	if !installed["extended"] {
		query.Set("extended", "1")
	}

	var (
		reply = responses.HistoryMessagesReply{}
		err   = responses.NewInternalError(bot.Call("messages.getHistory", query.Encode(), &reply))
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

func setHistoryMessagesStructProperty(query *url.Values, property types.HistoryMessages, installed map[string]bool) {
	if property.GroupID > 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))
	}

	if property.ChatID > 0 {
		query.Set("peer_id", fmt.Sprint(property.ChatID))

		installed["chatID"] = true
	} else if property.UserID > 0 {
		query.Set("user_id", fmt.Sprint(property.UserID))

		installed["chatID"] = true
	}

	if property.StartMessageID > 0 {
		query.Set("start_message_id", fmt.Sprint(property.StartMessageID))
	}

	if property.Offset > 0 {
		query.Set("offset", fmt.Sprint(property.Offset))

		installed["offset"] = true
	}

	if property.Count > 0 {
		query.Set("count", fmt.Sprint(property.Count))
	}

	if property.Rev {
		query.Set("rev", fmt.Sprint(conv.BooleanToInteger(property.Rev)))

		installed["rev"] = true
	}

	if property.Extended {
		query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property.Extended)))

		installed["extended"] = true
	}

	if len(property.Fields) > 0 {
		query.Set("fields", conv.ForParams(property.Fields))
	}
}

func setHistoryMessagesIntProperty(query *url.Values, property int, installed map[string]bool) {
	switch {
	case !installed["chatID"]:
		query.Set("peer_id", fmt.Sprint(property))

		installed["chatID"] = true
	case !installed["offset"]:
		query.Set("offset", fmt.Sprint(property))

		installed["offset"] = true
	default:
		query.Set("count", fmt.Sprint(property))
	}
}

func setHistoryMessagesBoolProperty(query *url.Values, property bool, installed map[string]bool) {
	if !installed["rev"] {
		query.Set("rev", fmt.Sprint(conv.BooleanToInteger(property)))

		installed["rev"] = true
	} else if !installed["extended"] {
		query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property)))

		installed["extended"] = true
	}
}
