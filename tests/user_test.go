package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	bot := NewBot()

	t.Run("getUser", func(t *testing.T) {
		user := bot.GetUser(1)
		assert.Equal(t, true, user.Error == nil, user.Error)
		assert.Equal(t, "Павел", user.Name, user)
	})
}
