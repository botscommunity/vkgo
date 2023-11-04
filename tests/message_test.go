package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	var (
		bot           = NewUser()
		chatMessageID int
	)

	t.Run("sendMessage", func(t *testing.T) {
		sent := bot.SendMessage(bot.ID, "Hello, bot!")
		assert.Equal(t, true, sent.Error == nil, sent.Error)

		chatMessageID = sent.ChatMessageID
		assert.Equal(t, true, chatMessageID != 0, sent)
	})
}
