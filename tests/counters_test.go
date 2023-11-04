package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounters(t *testing.T) {
	t.Run("counters", func(t *testing.T) {
		bot := NewUser()
		_, err := bot.GetCounters()
		assert.Equal(t, true, err == nil, err)
	})
}
