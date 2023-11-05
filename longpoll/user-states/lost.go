package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func lost(bot *api.Bot, session *responses.LongPollServer, messages Messages, _ scene.Scenes) update.Updates {
	if server := bot.GetUserLongPollServer(); server.Error == nil {
		session.Key = server.Key
		session.TS = server.TS
	}

	return update.Updates{
		TS: messages.TS,
	}
}
