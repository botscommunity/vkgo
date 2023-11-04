package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func getNameCase(fullCase string) string {
	cases := map[string]string{
		"nominative":   "nom",
		"genitive":     "gen",
		"dative":       "dat",
		"accusative":   "acc",
		"instrumental": "ins",
		"ablative":     "abl",
	}

	cas := cases[fullCase]
	if cas != "" {
		return cas
	}

	return fullCase
}

func (bot *Bot) GetUser(properties ...any) (user responses.User) {
	users, err := bot.GetUsers(properties...)
	if len(users) > 0 {
		user = users[0]
	}

	if err != nil {
		user.Error = err
	}

	return
}

func (bot *Bot) GetUsers(properties ...any) ([]responses.User, *responses.Error) {
	query := make(url.Values)

	for _, property := range properties {
		switch property := property.(type) {
		case types.User:
			if property.UserID != 0 {
				query.Set("user_ids", fmt.Sprint(property.UserID))
			} else if len(property.UserIDs) > 0 {
				query.Set("user_ids", conv.ForParams(property.UserIDs))
			}

			if property.NameCase != "" {
				if nameCase := getNameCase(property.NameCase); nameCase != "" {
					query.Set("name_case", nameCase)
				} else {
					query.Set("name_case", property.NameCase)
				}
			}

			if len(property.Fields) > 0 {
				query.Set("fields", conv.ForParams(property.Fields))
			}
		case int:
			query.Set("user_ids", fmt.Sprint(property))
		case []int:
			query.Set("user_ids", conv.ForParams(property))
		case string:
			if nameCase := getNameCase(property); nameCase != "" {
				query.Set("name_case", nameCase)
			} else {
				query.Set("name_case", property)
			}
		case []string:
			query.Set("fields", conv.ForParams(property))
		}
	}

	reply := responses.Users{}
	err := responses.NewInternalError(bot.Call("users.get", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Users, err
}
