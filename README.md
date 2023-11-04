<img src="preview.png" alt="Bots Community">

<p align="center">Flexible and high-performance VK API module Go</p>

<div align="center">
	<a href="https://pkg.go.dev/github.com/botscommunity/vkgo">
		<img src="https://img.shields.io/static/v1?label=go&message=reference&color=00add8&logo=go" />
	</a>
	<a href="http://www.opensource.org/licenses/MIT">
		<img src="https://img.shields.io/badge/license-MIT-fuchsia.svg" />
	</a>
    <a href="https://goreportcard.com/report/github.com/botscommunity/vkgo">
		<img src="https://goreportcard.com/badge/github.com/botscommunity/vkgo" />
	</a>
</div>


<h2 align="center">Instalation</h2>

```bash
go get github.com/botscommunity/vkgo
```

<h2 align="center">Getting started</h2>

Examples of working bots can be seen in the catalog [/samples](./samples)

A simple example of a LongPoll API Bot:

```go
package main

import (
  "github.com/botscommunity/vkgo/api"
  "github.com/botscommunity/vkgo/longpoll"
  "github.com/botscommunity/vkgo/scene"
  "github.com/botscommunity/vkgo/update"
  "os"
)

func main() {
  bot := api.New(os.Getenv("TOKEN"))

  messageScene := scene.Message(func(bot *api.Bot, message update.Message) {
    bot.SendMessage(message.ChatID, "echo message: "+message.Text)
  })

  longpoll.Start(bot, messageScene)
}
```

<h2 align="center">Help in solving problems</h2>

Don't know how to solve your problem? Ask the programmers from [our community](./CONTACTS.md). There is a chance that they have already dealt with this problem and are ready to help you