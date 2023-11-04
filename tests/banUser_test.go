package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBanUser(t *testing.T) {
	bot := NewUser()

	t.Run("banUser", func(t *testing.T) {
		ok, err := bot.BanUser(GetGroupID(), 1)
		assert.Equal(t, true, err == nil, err)
		assert.Equal(t, true, ok, err)
	})

	t.Run("getBannedUser", func(t *testing.T) {
		count, _, err := bot.GetBannedUsers(GetGroupID(), 1)
		assert.Equal(t, true, err == nil, err)
		assert.Equal(t, true, count != 0, err)
	})

	t.Run("unbanUser", func(t *testing.T) {
		ok, err := bot.UnbanUser(GetGroupID(), 1)
		assert.Equal(t, true, err == nil, err)
		assert.Equal(t, true, ok, err)
	})
}
