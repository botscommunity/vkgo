package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func expired(bot *api.Bot, session *responses.LongPollServer, _ update.Updates, _ scene.Scenes) {
	if server := bot.GetGroupLongPollServer(bot.ID); server.Error == nil {
		session.Key = server.Key
	}
}
