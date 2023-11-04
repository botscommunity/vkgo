package api

import (
	"fmt"
	"net/url"

	"github.com/botscommunity/vkgo/internal/pkg/conv"
	"github.com/botscommunity/vkgo/internal/pkg/responses"
	"github.com/botscommunity/vkgo/types"
)

func (bot *Bot) CreateChat(properties ...any) (chatID int, users []int, err *responses.Error) {
	query, userID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.CreateChat:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))
			}

			if property.UserID != 0 {
				query.Set("user_ids", fmt.Sprint(property.UserID))

				userID = true
			} else if len(property.UsersID) != 0 {
				query.Set("user_ids", conv.ForParams(property.UsersID))

				userID = true
			}

			if property.Name != "" {
				query.Set("title", property.Name)
			} else if property.Title != "" {
				query.Set("title", property.Title)
			}
		case int:
			if !userID {
				query.Set("user_ids", fmt.Sprint(property))

				userID = true
			} else {
				query.Set("group_id", fmt.Sprint(property))
			}
		case []int:
			query.Set("user_ids", conv.ForParams(property))

			userID = true
		case string:
			query.Set("title", property)
		}
	}

	reply := responses.CreateChatReply{}
	err = responses.NewInternalError(bot.Call("messages.createChat", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.ChatID, reply.Response.Users, err
}

func (bot *Bot) GetChat(properties ...any) (chat responses.Chat) {
	chats := bot.GetChats(properties...)

	if len(chats.Chats) > 0 {
		chat = chats.Chats[0]
	}

	if chats.Error != nil {
		chat.Error = chats.Error
	}

	return
}

func (bot *Bot) GetChats(properties ...any) responses.Chats {
	query, chatID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.Chat:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))
			}

			if property.ChatID != 0 {
				query.Set("peer_ids", fmt.Sprint(property.ChatID))

				chatID = true
			} else if len(property.ChatsID) != 0 {
				query.Set("peer_ids", conv.ForParams(property.ChatsID))

				chatID = true
			}

			if property.Extended {
				query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property.Extended)))
			}

			if len(property.Fields) != 0 {
				query.Set("fields", conv.ForParams(property.Fields))
			}
		case int:
			if !chatID {
				query.Set("peer_ids", fmt.Sprint(property))

				chatID = true
			} else {
				query.Set("group_id", fmt.Sprint(property))
			}
		case []int:
			query.Set("peer_ids", conv.ForParams(property))

			chatID = true
		case []string:
			query.Set("fields", conv.ForParams(property))
		case bool:
			query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	var (
		reply = responses.ChatReply{}
		err   = responses.NewInternalError(bot.Call("messages.getConversationsById", query.Encode(), &reply))
	)

	if err != nil {
		return responses.Chats{Error: err}
	}

	if reply.Error != nil {
		reply.Response.Error = reply.Error
	}

	return reply.Response
}

func (bot *Bot) GetChatLink(properties ...any) (link string, err *responses.Error) {
	query, chatID := make(url.Values), false

	for _, property := range properties {
		switch property := property.(type) {
		case types.ChatLink:
			if property.GroupID != 0 {
				query.Set("group_id", fmt.Sprint(property.GroupID))
			}

			if property.ChatID != 0 {
				query.Set("peer_id", fmt.Sprint(property.ChatID))

				chatID = true
			}

			if property.Reset {
				query.Set("reset", fmt.Sprint(conv.BooleanToInteger(property.Reset)))
			}
		case int:
			if !chatID {
				query.Set("peer_id", fmt.Sprint(property))

				chatID = true
			} else {
				query.Set("group_id", fmt.Sprint(property))
			}
		case bool:
			query.Set("reset", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	reply := responses.ChatLinkReply{}
	err = responses.NewInternalError(bot.Call("messages.getInviteLink", query.Encode(), &reply))

	if reply.Error != nil {
		err = reply.Error
	}

	return reply.Response.Link, err
}

func (bot *Bot) GetChatMembers(properties ...any) responses.ChatMembers {
	query, installed := make(url.Values), make(map[string]bool)

	for _, property := range properties {
		switch property := property.(type) {
		case types.ChatMembers:
			setChatMembersStructProperty(&query, property, installed)
		case int:
			setChatMembersIntProperty(&query, property, installed)
		case []string:
			query.Set("fields", conv.ForParams(property))
		case bool:
			query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property)))
		}
	}

	var (
		reply = responses.ChatMembersReply{}
		err   = responses.NewInternalError(bot.Call("messages.getConversationMembers", query.Encode(), &reply))
	)

	if err != nil {
		return responses.ChatMembers{Error: err}
	}

	if reply.Error != nil {
		reply.Response.Error = reply.Error
	}

	return reply.Response
}

func setChatMembersStructProperty(query *url.Values, property types.ChatMembers, installed map[string]bool) {
	if property.GroupID != 0 {
		query.Set("group_id", fmt.Sprint(property.GroupID))
	}

	if property.ChatID != 0 {
		query.Set("peer_id", fmt.Sprint(property.ChatID))

		installed["chatID"] = true
	}

	if property.Offset != 0 {
		query.Set("offset", fmt.Sprint(property.Offset))

		installed["offset"] = true
	}

	if property.Count != 0 {
		query.Set("count", fmt.Sprint(property.Count))
	}

	if property.Extended {
		query.Set("extended", fmt.Sprint(conv.BooleanToInteger(property.Extended)))
	}

	if len(property.Fields) > 0 {
		query.Set("fields", conv.ForParams(property.Fields))
	}
}

func setChatMembersIntProperty(query *url.Values, property int, installed map[string]bool) {
	switch {
	case !installed["chatID"]:
		query.Set("peer_id", fmt.Sprint(property))

		installed["chatID"] = true
	case !installed["offset"]:
		query.Set("offset", fmt.Sprint(property))

		installed["offset"] = true
	default:
		query.Set("group_id", fmt.Sprint(property))
	}
}
