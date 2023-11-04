package multibot

import (
	"fmt"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/longpoll"
	"github.com/botscommunity/vkgo/scene"
)

type MultiBot struct {
	Bots   []*api.Bot
	Scenes *scene.Scenes
}

func New(properties ...any) *MultiBot {
	mb := new(MultiBot)

	for _, property := range properties {
		switch property := property.(type) {
		case *api.Bot:
			mb.Bots = append(mb.Bots, property)
		case []*api.Bot:
			mb.Bots = append(mb.Bots, property...)
		case *scene.Scenes:
			mb.Scenes = property
		}
	}

	if mb.Bots == nil {
		mb.Bots = make([]*api.Bot, 0)
	}

	return mb
}

func (mb *MultiBot) Start() []*responses.Error {
	var (
		errors   = make([]*responses.Error, 0, len(mb.Bots))
		errorsCh = make(chan *responses.Error)
	)

	for _, bot := range mb.Bots {
		go listenBot(bot, mb.Scenes, errorsCh)
	}

	for err := range errorsCh {
		errors = append(errors, err)
	}

	return errors
}

func listenBot(bot *api.Bot, scenes *scene.Scenes, errors chan *responses.Error) {
	if bot.Logger != nil {
		bot.Logger.Info(fmt.Sprintf("MultiBot LongPoll  %d running", bot.ID))
	}

	if err := longpoll.Start(bot, scenes); err != nil {
		bot.Logger.Error(fmt.Sprintf("MultiBot LongPoll %d error is %s", bot.ID, err.Message))
		errors <- err
	} else {
		errors <- nil
	}
}
