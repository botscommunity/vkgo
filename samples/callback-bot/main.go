package main

import (
	"log"
	"os"

	"github.com/botscommunity/vkgo/api"
	"github.com/botscommunity/vkgo/callback"
	"github.com/botscommunity/vkgo/samples/scenes"
	"github.com/botscommunity/vkgo/scene"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	bot, err := api.New(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	messageScene := scene.Message(scenes.NewMessageScene())

	cb := callback.New(messageScene)
	cb.Add(bot)

	app := fiber.New()
	app.Post("/vk/:confirmation", cb.Router.Fiber)
	log.Fatalln(app.Listen(":3200"))
}
