package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_floatStone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		floatStone, err := NewFloatStone()
		if err != nil {
			t.Errorf("Error creating float stone: %v", err)
		}
		statBoost := floatStone.GetStatBoosts()
		assert.Equal(t, stats.Stats{Attack: 24.0}, statBoost)
	})
}
