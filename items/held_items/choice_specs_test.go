package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_choiceSpecs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		choiceSpecs, err := NewChoiceSpecs()
		if err != nil {
			t.Errorf("Error creating choice specs: %v", err)
		}
		statBoost := choiceSpecs.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{SpecialAttack: 39}, statBoost)
	})
}
