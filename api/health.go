package api

import (
	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) GetHealth() (statuses []responses.HealthStatus, err *responses.Error) {
	reply := responses.HealthReply{}
	err = responses.NewInternalError(bot.Call("documentation.getPlatformHealthStatuses", "", &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return statuses, err
}
