package main

import (
	"os"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/multibot"
	"github.com/botscommunity/vkgo/samples/scenes"
	"github.com/botscommunity/vkgo/scene"
	"go.uber.org/zap"
)

func main() {
	diamond, err := api.New(os.Getenv("DIAMOND_TOKEN"))
	if err != nil {
		panic(err)
	}

	sapphire, err := api.New(os.Getenv("SAPPHIRE_TOKEN"))
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	diamond.Logger = logger
	sapphire.Logger = logger

	bots := []*api.Bot{diamond, sapphire}

	messageScene := scene.Message(scenes.NewMessageScene())

	multibot.New(bots, messageScene, logger).Start()
}
