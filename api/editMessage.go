package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/keyboard"
	"github.com/botscommunity/vkgo/template"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) EditMessage(properties ...any) (ok bool, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.EditMessage:
			setEditMessageStructProperty(&query, property, installed)
		case int:
			setEditMessageIntProperty(&query, property, installed)
		case string:
			setEditMessageStringProperty(&query, property, installed)
		case []string:
			query.Set("attachment", conv.ForParams(property))
		case *template.Template:
			query.Set("template", property.JSON())
		case *keyboard.ReplyKeyboard:
			query.Set("keyboard", property.JSON())
		}
	}

	reply := responses.EditMessage{}
	err = responses.NewInternalError(bot.Call("messages.edit", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func setEditMessageStructProperty(query *url.Values, property types.EditMessage, installed map[string]bool) {
	if property.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))
	}

	if property.ChatID != 0 {
		query.Set("peer_id", fmt.Sprint(property.ChatID))

		installed["chatID"] = true
	}

	if property.Text != "" {
		query.Set("message", property.Text)

		installed["text"] = true
	}

	if property.MessageID != 0 {
		query.Set("message_id", fmt.Sprint(property.MessageID))
	} else if property.ChatMessageID != 0 {
		query.Set("conversation_message_id", fmt.Sprint(property.ChatMessageID))
	}

	if property.Attachment != "" {
		query.Set("attachment", property.Attachment)
	} else if len(property.Attachments) > 0 {
		query.Set("attachment", conv.ForParams(property.Attachments))
	}

	if property.Template != "" {
		query.Set("template", property.Template)
	} else if property.Keyboard != "" {
		query.Set("keyboard", property.Keyboard)
	}

	if property.KeepForwardMessages {
		query.Set("keep_forward_messages", fmt.Sprint(conv.BooleanToInteger(property.KeepForwardMessages)))
	}

	if property.KeepSnippets {
		query.Set("keep_snippets", fmt.Sprint(conv.BooleanToInteger(property.KeepSnippets)))
	}

	if property.DontParseLinks {
		query.Set("dont_parse_links", fmt.Sprint(conv.BooleanToInteger(property.DontParseLinks)))
	}

	if property.DisableMentions {
		query.Set("disable_mentions", fmt.Sprint(conv.BooleanToInteger(property.DisableMentions)))
	}

	if property.Latitude != "" {
		query.Set("lat", property.Latitude)
	}

	if property.Longitude != "" {
		query.Set("long", property.Longitude)
	}
}

func setEditMessageIntProperty(query *url.Values, property int, installed map[string]bool) {
	if !installed["chatID"] {
		query.Set("peer_id", fmt.Sprint(property))

		installed["chatID"] = true
	} else {
		query.Set("conversation_message_id", fmt.Sprint(property))
	}
}

func setEditMessageStringProperty(query *url.Values, property string, installed map[string]bool) {
	if !installed["text"] {
		query.Set("message", property)

		installed["text"] = true
	} else {
		query.Set("attachment", property)
	}
}
