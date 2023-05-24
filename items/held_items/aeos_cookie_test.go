package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_aeosCookie(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stacks := 5
		aeosCookie, err := NewAeosCookie(stacks)
		if err != nil {
			t.Errorf("Error creating aeos cookie: %v", err)
		}
		if aeosCookie.numStacks != stacks {
			t.Errorf("Expected aeos cookie stacks to be %d, got %v", stacks, aeosCookie.numStacks)
		}
		statBoost := aeosCookie.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Hp: 240.0 + (200.0 * float64(stacks))}, statBoost)
	})
}
