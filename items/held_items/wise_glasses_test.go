package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wiseGlasses(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		wiseGlasses, err := NewWiseGlasses(stats.Stats{SpecialAttack: 100.0})
		if err != nil {
			t.Errorf("Error creating wise glasses: %v", err)
		}
		statBoost := wiseGlasses.GetStatBoosts()
		assert.Equal(t, stats.Stats{SpecialAttack: 39.0 + (100.0 * 0.07)}, statBoost)
	})
}
