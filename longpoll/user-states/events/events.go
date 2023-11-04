package events

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/scene"
)

type Event func(_ *api.Bot, _ []any, _ scene.Scenes) string

func GetEvent(code float64) Event {
	events := map[float64]Event{
		4: message,
	}

	return events[code]
}
