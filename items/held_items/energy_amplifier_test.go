package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_energyAmplifier(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		energyAmplifier, err := NewEnergyAmplifier()
		if err != nil {
			t.Errorf("Error creating energy amplifier: %v", err)
		}
		statBoost := energyAmplifier.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{EnergyRate: 0.06, CooldownReduction: 0.045}, statBoost)
	})
}
