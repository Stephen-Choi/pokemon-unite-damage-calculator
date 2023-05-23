package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shellBell(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		shellBell, err := NewShellBell()
		if err != nil {
			t.Errorf("Error creating shell bell: %v", err)
		}
		statBoost := shellBell.GetStatBoosts()
		assert.Equal(t, stats.Stats{SpecialAttack: 24.0, CooldownReduction: 0.045}, statBoost)
	})
}
