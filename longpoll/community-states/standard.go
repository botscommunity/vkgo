package states

import (
	"sync"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
	"github.com/botscommunity/vkgo/update"
)

func standard(bot *api.Bot, session *responses.LongPollServer, updates update.Updates, scenes scene.Scenes) {
	wg := sync.WaitGroup{}
	for _, u := range updates.Updates {
		wg.Add(1)

		go func(u update.Update) {
			scene.Use(bot, u, scenes)
			wg.Done()
		}(u)
	}

	wg.Wait()

	session.TS = updates.TS
}
