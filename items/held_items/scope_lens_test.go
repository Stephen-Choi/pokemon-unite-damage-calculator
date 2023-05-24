package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_scopeLens(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		scopeLens, err := NewScopeLens()
		if err != nil {
			t.Errorf("Error creating scope lens: %v", err)
		}
		statBoost := scopeLens.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{CriticalHitChance: 0.06, CriticalHitDamage: 0.12}, statBoost)
	})
}
