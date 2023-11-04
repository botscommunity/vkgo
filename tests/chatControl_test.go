package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatControl(t *testing.T) {
	var (
		bot    = NewUser()
		chatID int
	)

	t.Run("createChat", func(t *testing.T) {
		createdChatID, _, err := bot.CreateChat("Chat by GitHub Actions")
		chatID = createdChatID

		assert.Equal(t, true, err == nil, err)
		assert.Equal(t, true, chatID != 0, err)
	})

	t.Run("getChat", func(t *testing.T) {
		chat := bot.GetChat(chatID)
		assert.Equal(t, true, chat.Error == nil, chat.Error)
	})

	t.Run("getChatLink", func(t *testing.T) {
		link, err := bot.GetChatLink(2e9 + chatID)
		assert.Equal(t, true, err == nil, err)
		assert.Equal(t, true, link != "", err)
	})

	t.Run("getChatMembers", func(t *testing.T) {
		members := bot.GetChatMembers(chatID)
		assert.Equal(t, true, members.Error == nil, members.Error)
	})
}
