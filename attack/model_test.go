package attack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetFramesDelayForAttackSpeed(t *testing.T) {
	t.Run("Default 0 attack speed", func(t *testing.T) {
		framesDelay := GetFramesDelayForAttackSpeed(0.0)
		assert.Equal(t, 60, framesDelay)
	})
	t.Run("Mid attack speed", func(t *testing.T) {
		framesDelay := GetFramesDelayForAttackSpeed(52)
		assert.Equal(t, 40, framesDelay)
	})
	t.Run("Max attack speed", func(t *testing.T) {
		framesDelay := GetFramesDelayForAttackSpeed(400)
		assert.Equal(t, 16, framesDelay)
	})
}
