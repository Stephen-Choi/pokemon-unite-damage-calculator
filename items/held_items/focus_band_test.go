package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_focusBand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		focusBand, err := NewFocusBand()
		if err != nil {
			t.Errorf("Error creating focus band: %v", err)
		}
		statBoost := focusBand.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Defense: 30, SpecialDefense: 30}, statBoost)
	})
}
