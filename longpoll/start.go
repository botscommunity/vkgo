package longpoll

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	communityStates "github.com/botscommunity/vkgo/longpoll/community-states"
	userStates "github.com/botscommunity/vkgo/longpoll/user-states"
	"github.com/botscommunity/vkgo/update"
)

func Start(properties ...any) *responses.Error {
	lp, err := New(properties...)
	if err != nil {
		return err
	}

	return lp.Start()
}

func (lp *Longpoll) Start() *responses.Error {
	if lp.Bot == nil {
		return responses.NewInternalError(ErrSessionClosed)
	}

	if lp.Session.Error != nil {
		return lp.Session.Error
	}

	updates := update.Updates{}

	switch lp.mode {
	case communityBot:
		for {
			if _, err := communityUpdates(lp, updates); err != nil {
				return err
			}
		}
	case userBot:
		for {
			if _, err := userUpdates(lp, updates); err != nil {
				return err
			}
		}
	}

	return nil
}

func communityUpdates(lp *Longpoll, updates update.Updates) (update.Updates, *responses.Error) {
	if lp.Session.Server == "" {
		return update.Updates{}, responses.NewInternalError(ErrSessionClosed)
	}

	link := fmt.Sprintf("%s?act=a_check&key=%s&ts=%v&wait=%d&mode=3&version=%f", lp.Session.Server, lp.Session.Key, lp.Session.TS, lp.Timeout, lp.Bot.Version)
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)

	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll NewRequest error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	response, err := lp.Bot.HTTPClient.Do(request)
	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll Do error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		if err := response.Body.Close(); err != nil {
			if lp.Bot.Logger != nil {
				lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll BodyClose error is %s", err.Error()))
			}
		}

		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll ReadAll error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	if err := json.Unmarshal(body, &updates); err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("GroupLongPoll JSON unmarshal error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	if lp.Bot.Logger != nil {
		lp.Bot.Logger.Info(fmt.Sprintf("GroupLongPoll response is %s", string(body)))
	}

	communityStates.ExecuteState(lp.Bot, lp.Session, updates, *lp.Scenes)

	return updates, responses.NewInternalError(err)
}

func userUpdates(lp *Longpoll, updates update.Updates) (update.Updates, *responses.Error) {
	if lp.Session.Server == "" {
		return update.Updates{}, responses.NewInternalError(ErrSessionClosed)
	}

	sessionTS := defineSessionTS(lp.Session.TS)

	link := fmt.Sprintf("https://%s?act=a_check&key=%s&ts=%+v&wait=%d&mode=3&version=%d", lp.Session.Server, lp.Session.Key, sessionTS, lp.Timeout, lp.Session.Version)
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)

	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll NewRequest error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	response, err := lp.Bot.HTTPClient.Do(request)
	if err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll Do error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		if err := response.Body.Close(); err != nil {
			if lp.Bot.Logger != nil {
				lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll BodyClose error is %s", err.Error()))
			}
		}

		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll ReadAll error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	var messages userStates.Messages
	if err := json.Unmarshal(body, &messages); err != nil {
		if lp.Bot.Logger != nil {
			lp.Bot.Logger.Error(fmt.Sprintf("UserLongPoll JSON unmarshal error is %s", err.Error()))
		}

		return updates, responses.NewInternalError(err)
	}

	if lp.Bot.Logger != nil {
		lp.Bot.Logger.Info(fmt.Sprintf("UserLongPoll response is %s", string(body)))
	}

	updates = userStates.ExecuteState(lp.Bot, lp.Session, messages, *lp.Scenes)

	return updates, responses.NewInternalError(err)
}

func defineSessionTS(sessionTS any) string {
	switch offset := sessionTS.(type) {
	case uint32, uint64:
		return fmt.Sprintf("%d", offset)
	case float32, float64:
		return fmt.Sprintf("%f", offset)
	}

	return ""
}
