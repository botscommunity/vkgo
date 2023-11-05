package main

import (
	"log"
	"os"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/longpoll"
	"github.com/botscommunity/vkgo/samples/scenes"
	"github.com/botscommunity/vkgo/scene"
)

func main() {
	bot, err := api.New(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	messageScene := scene.Message(scenes.NewMessageScene())

	lp, errLp := longpoll.New(bot, messageScene)
	if errLp != nil {
		panic(errLp)
	}

	log.Fatalln(lp.Start())
}
