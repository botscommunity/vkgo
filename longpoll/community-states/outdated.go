package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func outdated(_ *api.Bot, session *responses.LongPollServer, updates update.Updates, _ scene.Scenes) {
	session.TS = updates.TS
}
