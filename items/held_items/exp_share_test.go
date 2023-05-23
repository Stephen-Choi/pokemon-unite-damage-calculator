package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_scoreShield(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		scoreShield, err := NewScoreShield()
		if err != nil {
			t.Errorf("Error creating score shield: %v", err)
		}
		statBoost := scoreShield.GetStatBoosts()
		assert.Equal(t, stats.Stats{Hp: 450}, statBoost)
	})
}
