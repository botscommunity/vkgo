package api

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) SetGroupLongPollSettings(properties ...any) (ok bool, err *responses.Error) {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.SetGroupLongPoll:
			setSetGroupLongPollSettingsStructProperty(&query, property, installed)
		case int:
			query.Set("group_id", fmt.Sprint(property))

			installed["groupID"] = true
		case float32, float64:
			query.Set("api_version", strings.ReplaceAll(fmt.Sprintf("%f", property), "0", ""))

			installed["version"] = true
		case bool:
			query.Set("enabled", fmt.Sprintf("%t", property))

			installed["enabled"] = true
		}
	}

	if !installed["groupID"] && bot.IsGroup {
		query.Set("group_id", fmt.Sprint(bot.ID))
	}

	if !installed["version"] {
		query.Set("api_version", fmt.Sprintf("%f", bot.Version))
	}

	if !installed["enabled"] {
		query.Set("enabled", "true")
	}

	reply := responses.SetGroupLongPoll{}
	err = responses.NewInternalError(bot.Call("groups.setLongPollSettings", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	if reply.Code == 1 {
		ok = true
	}

	return ok, err
}

func setSetGroupLongPollSettingsStructProperty(query *url.Values, property types.SetGroupLongPoll, installed map[string]bool) {
	if property.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))

		installed["groupID"] = true
	}

	if property.Version != 0 {
		query.Set("api_version", fmt.Sprintf("%f", property.Version))

		installed["version"] = true
	}

	if property.Enabled {
		query.Set("enabled", fmt.Sprintf("%t", property.Enabled))

		installed["enabled"] = true
	}

	if property.Message {
		query.Set("message_new", fmt.Sprintf("%t", property.Message))
	}

	if property.ReplyMessage {
		query.Set("message_reply", fmt.Sprintf("%t", property.ReplyMessage))
	}

	if property.EditMessage {
		query.Set("message_edit", fmt.Sprintf("%t", property.EditMessage))
	}

	if property.EnableMessage {
		query.Set("message_allow", fmt.Sprintf("%t", property.EnableMessage))
	}

	if property.DisableMessage {
		query.Set("message_deny", fmt.Sprintf("%t", property.DisableMessage))
	}

	if property.Typing {
		query.Set("message_typing_state", fmt.Sprintf("%t", property.Typing))
	}

	if property.Callback {
		query.Set("message_event", fmt.Sprintf("%t", property.Callback))
	}
}
