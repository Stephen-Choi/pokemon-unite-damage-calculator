package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_weaknessPolicy(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		weaknessPolicy, err := NewWeaknessPolicy()
		if err != nil {
			t.Errorf("Error creating weakness policy: %v", err)
		}
		statBoost := weaknessPolicy.GetStatBoosts()
		assert.Equal(t, stats.Stats{Hp: 210.0, Attack: 15.0}, statBoost)
	})
}
