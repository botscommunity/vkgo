package api

import (
	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) GetCounters() (responses.Counters, *responses.Error) {
	reply := responses.CountersReply{}
	if err := responses.NewInternalError(bot.Call("account.getCounters", "", &reply)); err != nil {
		return responses.Counters{}, err
	}

	return reply.Response, reply.Error
}
