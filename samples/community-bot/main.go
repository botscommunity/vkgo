package main

import (
	"fmt"
	"os"

	"github.com/botscommunity/vkgo/api"
)

func main() {
	bot, err := api.New(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	users, usersErr := bot.GetUsers([]int{1, 2})
	if usersErr != nil {
		panic(usersErr)
	}

	fmt.Println(users)
}
