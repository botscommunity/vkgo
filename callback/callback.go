package callback

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/scene"
)

// Callback 🌍 is a configuration structure of bots, scenes, routers, and "everyone".
type Callback struct {
	Bots     map[int]*api.Bot
	Scenes   *scene.Scenes
	Router   Router
	everyone bool
}

// New 🌠 returns an instance of *Callback
// Accepts 🤖 *api.Bot or 🎪 *scene.Scenes arguments.
func New(properties ...any) *Callback {
	cb := &Callback{
		Bots:   make(map[int]*api.Bot),
		Scenes: scene.New(),
	}

	cb.Router = Router{
		Config: cb,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *api.Bot:
			cb.Bots[property.ID] = property
		case *scene.Scenes:
			cb.Scenes = property
		}
	}

	return cb
}
