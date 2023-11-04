package tests_test

import (
	"os"
	"strconv"

	"github.com/botscommunity/vkgo/api"
)

func NewUser() *api.Bot {
	bot, err := api.New(os.Getenv("USER_TOKEN"))
	if err != nil {
		panic(err)
	}

	return bot
}

func NewBot() *api.Bot {
	bot, err := api.New(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	return bot
}

func GetGroupID() int {
	var (
		env          = os.Getenv("GROUP_ID")
		groupID, err = strconv.Atoi(env)
	)

	if err != nil {
		panic(err)
	}

	return groupID
}
