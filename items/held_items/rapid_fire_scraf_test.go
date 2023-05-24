package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rapidFireScarf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rapidFireScarf, err := NewRapidFireScarf()
		if err != nil {
			t.Errorf("Error creating Rapid Fire Scarf: %v", err)
		}
		statBoost := rapidFireScarf.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Attack: 12.0, AttackSpeed: 0.09}, statBoost)
	})
}
