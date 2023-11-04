package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) PinMessage(properties ...any) responses.PinMessage {
	query, chatID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.PinMessage:
			if property.ChatID != 0 {
				query.Set("peer_id", fmt.Sprint(property.ChatID))

				chatID = true
			}

			if property.MessageID != 0 {
				query.Set("message_id", fmt.Sprint(property.MessageID))
			} else if property.ChatMessageID != 0 {
				query.Set("conversation_message_id", fmt.Sprint(property.ChatMessageID))
			}
		case int:
			if !chatID {
				query.Set("peer_id", fmt.Sprint(property))

				chatID = true
			} else {
				query.Set("conversation_message_id", fmt.Sprint(property))
			}
		}
	}

	var (
		pinned = responses.PinMessageReply{}
		err    = responses.NewInternalError(bot.Call("messages.pin", query.Encode(), &pinned))
	)

	if err != nil {
		pinned.Response.Error = err
		return pinned.Response
	}

	if pinned.Error != nil {
		pinned.Response.Error = pinned.Error
	}

	return pinned.Response
}
