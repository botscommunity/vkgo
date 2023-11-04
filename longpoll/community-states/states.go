package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

type state func(bot *api.Bot, session *responses.LongPollServer, updates update.Updates, scenes scene.Scenes)

func ExecuteState(bot *api.Bot, session *responses.LongPollServer, updates update.Updates, scenes scene.Scenes) {
	states := map[int]state{
		0: standard,
		1: outdated,
		2: expired,
		3: lost,
	}

	if state := states[updates.Failed]; state != nil {
		state(bot, session, updates, scenes)
	}
}
