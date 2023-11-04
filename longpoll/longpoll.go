package longpoll

import (
	"errors"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/scene"
)

const DefaultTimeout = 60

var (
	ErrNotFound      = errors.New("bot not found")
	ErrSessionClosed = errors.New("session closed")
)

const (
	communityBot = 1
	userBot      = 2
)

// Longpoll - basic configuration structure
// Bot - pointer to the bot parameters
// Session - pointer to the session, received from the Create function
// Scenes - pointer to the list of scenes to process the types of events
// Timeout - time of waiting for the next request.
type Longpoll struct {
	Bot     *api.Bot
	Session *responses.LongPollServer
	Scenes  *scene.Scenes
	Timeout int
	mode    int // 1 - community, 2 - user
}

func New(properties ...any) (*Longpoll, *responses.Error) {
	lp := &Longpoll{
		Session: new(responses.LongPollServer),
		Scenes:  scene.New(),
		Timeout: DefaultTimeout,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *api.Bot:
			lp.Bot = property
		case *scene.Scenes:
			lp.Scenes = property
		}
	}

	if lp.Bot.ID == 0 {
		return lp, responses.NewInternalError(ErrNotFound)
	}

	switch lp.Bot.IsGroup {
	case true:
		server := lp.Bot.GetGroupLongPollServer()
		if lp.Session = &server; lp.Session.Error != nil {
			return lp, lp.Session.Error
		}

		lp.mode = communityBot
	default:
		server := lp.Bot.GetUserLongPollServer()
		if lp.Session = &server; lp.Session.Error != nil {
			return lp, lp.Session.Error
		}

		lp.mode = userBot
	}

	return lp, nil
}
