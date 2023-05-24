package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_muscleBand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		muscleBand, err := NewMuscleBand()
		if err != nil {
			t.Errorf("Error creating muscle band: %v", err)
		}
		statBoost := muscleBand.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Attack: 15, AttackSpeed: 0.075}, statBoost)
	})
}
