package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) GetStorageKeys(userID int) (keys []string, err *responses.Error) {
	query := make(url.Values)
	query.Set("user_id", fmt.Sprint(userID))

	reply := responses.StorageKeys{}
	err = responses.NewInternalError(bot.Call("storage.getKeys", query.Encode(), &keys))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response, err
}

func (bot *Bot) GetStorage(properties ...any) (storages responses.Storages, err *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			query.Set("key", property)
		case []string:
			query.Set("keys", conv.ForParams(property))
		case int:
			query.Set("user_id", fmt.Sprint(property))
		}
	}

	reply := responses.Storage{}
	err = responses.NewInternalError(bot.Call("storage.get", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response, err
}

func (bot *Bot) SetStorage(properties ...any) (result any, err *responses.Error) {
	query, key := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case string:
			if !key {
				query.Set("key", property)

				key = true
			} else {
				query.Set("value", property)
			}
		case int:
			query.Set("user_id", fmt.Sprint(property))
		}
	}

	reply := responses.SetStorage{}
	err = responses.NewInternalError(bot.Call("storage.set", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response, err
}
