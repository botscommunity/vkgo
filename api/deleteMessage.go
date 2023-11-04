package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) Spam(properties ...any) (message responses.DeleteMessage, err *responses.Error) {
	messages, err := bot.Spams(properties...)
	if len(messages) > 0 {
		message = messages[0]
	}

	return message, err
}

func (bot *Bot) Spams(properties ...any) (messages []responses.DeleteMessage, err *responses.Error) {
	properties = append(properties, types.DeleteMessage{Spam: true})
	return bot.DeleteMessages(properties...)
}

func (bot *Bot) DeleteMessage(properties ...any) (message responses.DeleteMessage, err *responses.Error) {
	messages, err := bot.DeleteMessages(properties...)
	if len(messages) > 0 {
		message = messages[0]
	}

	return message, err
}

func (bot *Bot) DeleteMessages(properties ...any) (messages []responses.DeleteMessage, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.DeleteMessage:
			setDeleteMessageStructProperty(&query, property, installed)
		case int:
			setDeleteMessageIntProperty(&query, property, installed)
		case []int:
			query.Set("cmids", conv.ForParams(property))

			installed["messageID"] = true
		case bool:
			setDeleteMessageBoolProperty(&query, property, installed)
		}
	}

	reply := responses.DeleteMessages{}
	err = responses.NewInternalError(bot.Call("messages.delete", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Messages, err
}

func setDeleteMessageStructProperty(query *url.Values, property types.DeleteMessage, installed map[string]bool) {
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
	} else if len(property.MessagesID) != 0 {
		query.Set("message_ids", conv.ForParams(property.MessagesID))

		installed["messageID"] = true
	}

	if property.ChatMessageID != 0 {
		query.Set("cmids", fmt.Sprint(property.ChatMessageID))

		installed["messageID"] = true
	} else if len(property.ChatMessagesID) != 0 {
		query.Set("cmids", conv.ForParams(property.ChatMessagesID))

		installed["messageID"] = true
	}

	if property.Spam {
		query.Set("spam", fmt.Sprint(conv.BooleanToInteger(property.Spam)))

		installed["spam"] = true
	}

	if property.Everyone {
		query.Set("delete_for_all", fmt.Sprint(conv.BooleanToInteger(property.Everyone)))

		installed["everyone"] = true
	}
}

func setDeleteMessageIntProperty(query *url.Values, property int, installed map[string]bool) {
	switch {
	case !installed["chatID"]:
		query.Set("peer_id", fmt.Sprint(property))

		installed["chatID"] = true
	case !installed["messageID"]:
		query.Set("cmids", fmt.Sprint(property))

		installed["messageID"] = true
	default:
		query.Set("delete_for_all", fmt.Sprint(property))
	}
}

func setDeleteMessageBoolProperty(query *url.Values, property bool, installed map[string]bool) {
	if !installed["everyone"] {
		query.Set("delete_for_all", fmt.Sprint(conv.BooleanToInteger(property)))

		installed["everyone"] = true
	} else if !installed["spam"] {
		query.Set("spam", fmt.Sprint(conv.BooleanToInteger(property)))

		installed["spam"] = true
	}
}
