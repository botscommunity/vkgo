package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func expired(bot *api.Bot, session *responses.LongPollServer, messages Messages, _ scene.Scenes) update.Updates {
	if server := bot.GetUserLongPollServer(); server.Error.Code == 0 {
		session.Key = server.Key
	}

	return update.Updates{TS: messages.TS}
}
