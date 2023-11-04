package states

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/longpoll/user-states/events"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func standard(bot *api.Bot, session *responses.LongPollServer, messages Messages, scenes scene.Scenes) update.Updates {
	updates := update.Updates{TS: messages.TS}

	for _, upd := range messages.Updates {
		go func(upd []any) {
			if len(upd) == 0 {
				return
			}

			if typ, ok := upd[0].(float64); ok {
				if event := events.GetEvent(typ); event != nil {
					event(bot, upd, scenes)
				}
			}
		}(upd)
	}

	session.TS = messages.TS

	return updates
}
