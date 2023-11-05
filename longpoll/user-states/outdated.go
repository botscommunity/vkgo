package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func outdated(_ *api.Bot, session *responses.LongPollServer, messages Messages, _ scene.Scenes) update.Updates {
	session.TS = messages.TS

	return update.Updates{
		TS: messages.TS,
	}
}
