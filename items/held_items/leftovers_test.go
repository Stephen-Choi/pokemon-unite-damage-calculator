package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_leftovers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		leftovers, err := NewLeftovers()
		if err != nil {
			t.Errorf("Error creating leftovers: %v", err)
		}
		statBoost := leftovers.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Hp: 360}, statBoost)
	})
}
