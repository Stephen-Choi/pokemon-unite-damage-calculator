package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_slickSpoon(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		slickSpoon, err := NewSlickSpoon()
		if err != nil {
			t.Errorf("Error creating Slick Spoon: %v", err)
		}
		statBoost := slickSpoon.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Hp: 210, SpecialAttack: 30}, statBoost)
	})
}
