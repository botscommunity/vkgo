package scenes

import (
	"fmt"
	"time"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func NewMessageScene() scene.MessageFunc {
	return func(bot *api.Bot, message update.Message) {
		if message.Text == "ping" {
			var (
				oldTime  = time.Now()
				pingText = "🚀 Pong!"
			)

			sent := bot.SendMessage(message.ChatID, pingText)
			if sent.Error == nil {
				var (
					durationText = fmt.Sprintf("🕒 Duration: %s", time.Since(oldTime))
					text         = fmt.Sprintf("%s\n%s", pingText, durationText)
				)

				if _, err := bot.EditMessage(sent.ChatID, sent.ChatMessageID, text); err != nil {
					panic(err)
				}
			}
		}
	}
}
