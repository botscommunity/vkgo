package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) Execute(properties ...any) (result any, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.Execute:
			setExecuteStructProperty(&query, property, installed)
		case string:
			query.Set("code", property)
		case float32, float64:
			query.Set("func_v", fmt.Sprintf("%f", property))

			installed["version"] = true
		}
	}

	if !installed["version"] {
		query.Set("func_v", fmt.Sprintf("%f", bot.Version))
	}

	reply := responses.Execute{}
	err = responses.NewInternalError(bot.Call("execute", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Result, err
}

func setExecuteStructProperty(query *url.Values, property types.Execute, installed map[string]bool) {
	if property, ok := property.Code.(string); ok {
		query.Set("code", property)
	}

	if property.Version != 0 {
		query.Set("func_v", fmt.Sprintf("%f", property))

		installed["version"] = true
	}
}
