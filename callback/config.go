package callback

import "github.com/botscommunity/vkgo/api"

func (cb *Callback) Add(bot *api.Bot) bool {
	cb.Bots[bot.ID] = bot
	return cb.Bots[bot.ID] != nil
}

func (cb *Callback) Remove(bot *api.Bot) bool {
	delete(cb.Bots, bot.ID)
	return cb.Bots[bot.ID] == nil
}
