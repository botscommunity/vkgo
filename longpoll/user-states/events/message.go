package events

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/types"
)

func message(bot *api.Bot, messages []any, scenes scene.Scenes) string {
	if len(messages) >= 1 {
		if typ, ok := messages[1].(float64); ok {
			if scenes.MessageFunc != nil {
				_, message, _, _, _ := bot.GetMessage(types.Message{
					MessageID: int(typ),
				})
				scenes.MessageFunc(bot, message)
			}
		}

		return "message_new"
	}

	return ""
}
