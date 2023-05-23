package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_drainCrown(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		drainCrown, err := NewDrainCrown()
		if err != nil {
			t.Errorf("Error creating drain crown: %v", err)
		}
		statBoost := drainCrown.GetStatBoosts()
		assert.Equal(t, stats.Stats{Hp: 120, Attack: 18}, statBoost)
	})
}
