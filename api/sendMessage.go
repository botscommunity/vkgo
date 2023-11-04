package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/keyboard"
	"github.com/botscommunity/vkgo/template"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) SendSticker(chatID, stickerID int, properties ...any) (message responses.SendMessage) {
	messages, err := bot.SendStickers(chatID, stickerID, properties...)

	if len(messages) > 0 {
		message = messages[0]
	}

	if err != nil {
		message.Error = err
	}

	return
}

func (bot *Bot) SendMessage(properties ...any) (message responses.SendMessage) {
	messages, err := bot.SendMessages(properties...)

	if len(messages) > 0 {
		message = messages[0]
	}

	if err != nil {
		message.Error = err
	}

	return
}

func (bot *Bot) SendStickers(chatID, stickerID int, properties ...any) ([]responses.SendMessage, *responses.Error) {
	properties = append(properties, types.SendMessage{
		ChatID:    chatID,
		StickerID: stickerID,
	})

	return bot.SendMessages(properties...)
}

func (bot *Bot) SendMessages(properties ...any) (messages []responses.SendMessage, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.SendMessage:
			setSendMessageStructProperty(&query, property, installed)
		case int:
			query.Set("peer_ids", fmt.Sprint(property))
		case []int:
			query.Set("peer_ids", conv.ForParams(property))
		case string:
			setSendMessageStringProperty(&query, property, installed)
		case []string:
			query.Set("attachment", conv.ForParams(property))
		case *template.Template:
			query.Set("template", property.JSON())
		case *keyboard.ReplyKeyboard:
			query.Set("keyboard", property.JSON())
		}
	}

	if !installed["randomID"] {
		query.Set("random_id", fmt.Sprint(types.Random()))
	}

	reply := responses.SendMessages{}
	err = responses.NewInternalError(bot.Call("messages.send", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Messages, err
}
func setSendMessageChatAndUserIDs(query *url.Values, property types.SendMessage) {
	if property.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))
	}

	if property.ChatID != 0 {
		query.Set("peer_ids", fmt.Sprint(property.ChatID))
	} else if property.ChatsID != nil {
		query.Set("peer_ids", conv.ForParams(property.ChatsID))
	}

	if property.UserID != 0 {
		query.Set("user_ids", fmt.Sprint(property.UserID))
	} else if len(property.UsersID) != 0 {
		query.Set("user_ids", conv.ForParams(property.UsersID))
	}
}

func setSendMessageContent(query *url.Values, property types.SendMessage, installed map[string]bool) {
	if property.Text != "" {
		query.Set("message", property.Text)

		installed["text"] = true
	}

	if property.Attachment != "" {
		query.Set("attachment", property.Attachment)
	} else if property.Attachments != nil {
		query.Set("attachment", conv.ForParams(property.Attachments))
	}

	if property.StickerID != 0 {
		query.Set("sticker_id", fmt.Sprint(property.StickerID))

		installed["text"] = true
	}
}

func setSendMessageForwardingOptions(query *url.Values, property types.SendMessage) {
	if property.Forward.ChatID != 0 {
		forward, _ := json.Marshal(property.Forward)

		query.Set("forward", string(forward))
	}
}

func setSendMessageKeyboardAndTemplateOptions(query *url.Values, property types.SendMessage) {
	if property.Template != "" {
		query.Set("template", property.Template)
	} else if property.Keyboard != "" {
		query.Set("keyboard", property.Keyboard)
	}
}

func setSendMessageOtherOptions(query *url.Values, property types.SendMessage) {
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

	if property.Intent != "" {
		query.Set("intent", property.Intent)
	}

	if property.SubscribeID != 0 {
		query.Set("subscribe_id", fmt.Sprint(property.SubscribeID))
	}

	if property.ExpireTime != 0 {
		query.Set("expire_ttl", fmt.Sprint(property.ExpireTime))
	}
}

func setSendMessageStructProperty(query *url.Values, property types.SendMessage, installed map[string]bool) {
	setSendMessageChatAndUserIDs(query, property)
	setSendMessageContent(query, property, installed)
	setSendMessageForwardingOptions(query, property)
	setSendMessageKeyboardAndTemplateOptions(query, property)
	setSendMessageOtherOptions(query, property)
}

func setSendMessageStringProperty(query *url.Values, property string, installed map[string]bool) {
	if !installed["text"] {
		query.Set("message", property)

		installed["text"] = true
	} else {
		query.Set("attachment", property)
	}
}
