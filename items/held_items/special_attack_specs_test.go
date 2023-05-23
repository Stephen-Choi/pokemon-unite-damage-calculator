package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_specialAttackSpecs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stacks := 6
		specialAttackSpecs, err := NewSpecialAttackSpecs(stacks)
		if err != nil {
			t.Errorf("Error creating special attack specs: %v", err)
		}
		if specialAttackSpecs.numStacks != stacks {
			t.Errorf("Expected special attack specs stacks to be %d, got %v", stacks, specialAttackSpecs.numStacks)
		}
		statBoost := specialAttackSpecs.GetStatBoosts()
		assert.Equal(t, stats.Stats{SpecialAttack: 24.0 + (16.0 * float64(stacks))}, statBoost)
	})
}
