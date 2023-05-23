package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_attackWeight(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stacks := 2
		attackWeight, err := NewAttackWeight(stacks)
		if err != nil {
			t.Errorf("Error creating attack weight: %v", err)
		}
		if attackWeight.numStacks != stacks {
			t.Errorf("Expected attack weight stacks to be %d, got %v", stacks, attackWeight.numStacks)
		}
		statBoost := attackWeight.GetStatBoosts()
		assert.Equal(t, stats.Stats{Attack: 18.0 + (12.0 * float64(stacks))}, statBoost)
	})
}
