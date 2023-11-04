package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
	"github.com/botscommunity/vkgo/update"
)

func (bot *Bot) GetMessage(properties ...any) (count int, message update.Message, groups []responses.Group, users []responses.User, err *responses.Error) {
	messages := bot.GetMessages(properties...)
	count = messages.Count
	groups = messages.Groups
	users = messages.Profiles
	err = messages.Error

	if len(messages.Messages) > 0 {
		message = messages.Messages[0]
	}

	return
}

func (bot *Bot) GetMessages(properties ...any) responses.Messages {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.Message:
			setMessageStructProperty(&query, property, installed)
		case int:
			setMessageIntProperty(&query, property, installed)
		case []int:
			query.Set("cmids", conv.ForParams(property))

			installed["messageID"] = true
		case []string:
			query.Set("fields", conv.ForParams(property))
		case bool:
			query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	var (
		messages = responses.MessagesReply{}
		err      = bot.Call("messages.getById", query.Encode(), &messages)
	)

	if err != nil {
		messages.Response.Error = responses.NewInternalError(err)
		return messages.Response
	}

	decodedMessages, err := messages.Response.Decode()
	if err != nil {
		messages.Response.Error = responses.NewInternalError(err)
		return messages.Response
	}

	messages.Response = decodedMessages
	if messages.Error != nil {
		messages.Response.Error = messages.Error
	}

	return messages.Response
}

func setMessageStructProperty(query *url.Values, property types.Message, installed map[string]bool) {
	if property.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))
	}

	if property.ChatID != 0 {
		query.Set("peer_id", fmt.Sprint(property.ChatID))

		installed["chatID"] = true
	}

	if property.MessageID != 0 {
		query.Set("message_ids", fmt.Sprint(property.MessageID))

		installed["messageID"] = true
	} else if property.ChatMessageID != 0 {
		query.Set("cmids", fmt.Sprint(property.ChatMessageID))

		installed["messageID"] = true
	}

	if property.MessagesID != nil {
		query.Set("message_ids", conv.ForParams(property.MessagesID))

		installed["messageID"] = true
	} else if property.ChatMessagesID != nil {
		query.Set("cmids", conv.ForParams(property.ChatMessagesID))

		installed["messageID"] = true
	}

	if property.Fields != nil {
		query.Set("fields", conv.ForParams(property.Fields))
	}

	if property.Extended {
		query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property.Extended)))
	}

	if property.Characters != 0 {
		query.Set("preview_length", fmt.Sprint(property.Characters))
	}
}

func setMessageIntProperty(query *url.Values, property int, installed map[string]bool) {
	switch {
	case !installed["chatID"]:
		query.Set("peer_id", fmt.Sprint(property))

		installed["chatID"] = true
	case !installed["messageID"]:
		query.Set("cmids", fmt.Sprint(property))

		installed["messageID"] = true
	default:
		query.Set("group_id", fmt.Sprint(property))
	}
}
