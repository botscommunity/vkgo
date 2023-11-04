package tests_test

import (
	"testing"

	"github.com/botscommunity/vkgo/types"
	"github.com/stretchr/testify/assert"
)

func Test_Ban(t *testing.T) {
	bot := NewUser()

	t.Run("ban", func(t *testing.T) {
		_, err := bot.Ban(1)
		assert.Equal(t, true, err == nil, err)
	})

	t.Run("getBanned", func(t *testing.T) {
		_, _, _, err := bot.GetBanned(types.Banned{Count: 0})
		assert.Equal(t, true, err == nil, err)
	})

	t.Run("unban", func(t *testing.T) {
		_, err := bot.Unban(1)
		assert.Equal(t, true, err == nil, err)
	})
}
