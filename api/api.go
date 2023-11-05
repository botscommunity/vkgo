package api

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/botscommunity/vkgo/internal/pkg/consts"

	"github.com/botscommunity/botsgo"
	"go.uber.org/zap"
)

// Bot ✨ is an object that defines the properties of a bot
// 🏙️ The ID is the unique identifier of the bot. It contains either a community ID or a user page ID
// 🏡 A token is a key to access your application (bot)
// ☕ Version - VK API version, it is strongly recommended to set 5.154000
// 📺 IsGroup is set automatically, guided by the information about the created bot
// 🧢 Logger is a pointer to the Logger object of the zap library. Used during development to detect problems during operation
// 👆 ContentType a simpler  format of response data — https://en.wikipedia.org/wiki/MessagePack
// 💕 Limit - limiting the number of requests per second depending on the type of bot
// 🌿 errorHandler performs the task of transmitting an error for further processing.
type Bot struct {
	ID          int     `json:"id"`
	Token       string  `json:"token"`
	Version     float32 `json:"version"`
	IsGroup     bool    `json:"is_group"`
	Deprecated  Deprecated
	ContentType string
	Limit       int
	mutex       sync.Mutex
	time        time.Time
	rps         int
	*botsgo.Bot
}

func NewBot(token string, communityID int, properties ...any) (*Bot, error) {
	return New(&Bot{
		ID:    communityID,
		Token: token,
		Limit: consts.CommunityAPILimit,
	}, properties)
}

func NewUser(token string, userID int, properties ...any) (*Bot, error) {
	return New(&Bot{
		ID:    userID,
		Token: token,
		Limit: consts.CommunityAPILimit,
	}, properties)
}

// New 📺 creating a page/community bot
// 🏙️ String - token
// 🎮 Float32 - version
// 💬 Int - page/community ID
// 💗 Bot{} - advanced settings
// 🍭 Returns an instance of the *Bot.
func New(properties ...any) (*Bot, error) {
	client, err := botsgo.NewBot("https://api.vk.com/method")
	if err != nil {
		return nil, err
	}

	client.Logger = zap.NewNop()

	bot := &Bot{
		Version:     consts.DefaultAPIVersion,
		ContentType: "application/msgpack",
		Limit:       consts.UserAPILimit,
		Bot:         client,
	}

	for _, property := range properties {
		switch property := property.(type) {
		case *Bot:
			botProperties(bot, property)
		case int:
			bot.ID = property
		case string:
			bot.Token = property
		case float32:
			bot.Version = property
		case *zap.Logger:
			bot.Bot.Logger = property
		}
	}

	if bot.ID == 0 {
		bot.defineID()
	}

	bot.Deprecated.Bot = bot

	return bot, nil
}

func botProperties(bot, property *Bot) {
	if property.ID != 0 {
		bot.ID = property.ID
	}

	if property.Token != "" {
		bot.Token = property.Token
	}

	if property.Version != 0 {
		bot.Version = property.Version
	}

	if property.ContentType != "" {
		bot.ContentType = property.ContentType
	}

	if property.Limit != 0 {
		bot.Limit = property.Limit
	}

	if property.Bot != nil {
		if property.Bot.HTTPClient != nil {
			bot.Bot.HTTPClient = property.Bot.HTTPClient
		}

		if property.Bot.API != "" {
			bot.Bot.API = property.Bot.API
		}

		if property.Bot.Logger != nil {
			bot.Bot.Logger = property.Bot.Logger
		}
	}
}

// Call 🔥🌅 sending an API request
// 🌌❄ The first argument is the name of the method,
// 🌺🌼🌹 the second is the URL string,
// 🦋🌹🎲 then usually a pointer to the structure where the response will be written
// 🌍🌊🐠 May return an error.
func (bot *Bot) Call(method string, queryParams string, responseStruct any) error {
	// code from VK SDK, thanks <3
	if bot.Limit > 0 {
		bot.mutex.Lock()

		sleepTime := time.Second - time.Since(bot.time)
		if sleepTime < 0 {
			bot.time = time.Now()
			bot.rps = 0
		} else if bot.rps == bot.Limit {
			time.Sleep(sleepTime)
			bot.time = time.Now()
			bot.rps = 0
		}

		bot.rps++
		bot.mutex.Unlock()
	}

	if bot.ContentType == "application/msgpack" {
		method += ".msgpack"
	}

	return bot.Bot.Call(botsgo.CallOptions{
		Method:   http.MethodPost,
		Path:     fmt.Sprintf("%s?access_token=%s&v=%f&%s", method, bot.Token, bot.Version, queryParams),
		Response: responseStruct,
	})
}
func (bot *Bot) defineID() {
	if group := bot.GetGroup(); group.Error == nil {
		bot.ID = group.ID
		bot.IsGroup = true
		bot.Limit = consts.CommunityAPILimit
	} else if user := bot.GetUser(); user.Error == nil {
		bot.ID = user.ID
	}
}
