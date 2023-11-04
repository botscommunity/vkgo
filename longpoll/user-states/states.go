package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

type Messages struct {
	TS         any     `json:"ts,omitempty"`
	Updates    [][]any `json:"updates,omitempty"`
	Failed     int     `json:"failed,omitempty"`
	MaxVersion int     `json:"max_version"`
	MinVersion int     `json:"min_version"`
}

type state func(bot *api.Bot, session *responses.LongPollServer, messages Messages, scenes scene.Scenes) update.Updates

func ExecuteState(bot *api.Bot, session *responses.LongPollServer, messages Messages, scenes scene.Scenes) update.Updates {
	states := map[int]state{
		0: standard,
		1: outdated,
		2: expired,
		3: lost,
		4: invalid,
	}

	updates := update.Updates{}
	if state := states[updates.Failed]; state != nil {
		updates = state(bot, session, messages, scenes)
	}

	return updates
}
