package callback

import (
	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/scene"
)

// Callback - structure of the connected bot.
type Callback struct {
	Bots     map[int]*api.Bot
	Scenes   *scene.Scenes
	Router   Router
	everyone bool
}

// New - server creation
// Accepts the option pointer argument and optional scenes
// Returns server options.
func New(properties ...any) *Callback {
	cb := &Callback{
		Bots:   make(map[int]*api.Bot),
		Scenes: scene.New(),
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *api.Bot:
			cb.Bots[property.ID] = property
		case *scene.Scenes:
			cb.Scenes = property
		}
	}

	cb.Router = Router{
		Config: cb,
	}

	return cb
}
