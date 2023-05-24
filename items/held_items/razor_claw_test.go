package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_razorClaw(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		razorClaw, err := NewRazorClaw()
		if err != nil {
			t.Errorf("Error creating razor claw: %v", err)
		}
		statBoost := razorClaw.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Attack: 15, CriticalHitChance: 0.021}, statBoost)
	})
}
