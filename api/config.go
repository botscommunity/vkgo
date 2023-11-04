package api

func (bot *Bot) SetMessagePack() string {
	bot.ContentType = "application/msgpack"
	return bot.ContentType
}

func (bot *Bot) SetJSON() string {
	bot.ContentType = "application/json"
	return bot.ContentType
}
