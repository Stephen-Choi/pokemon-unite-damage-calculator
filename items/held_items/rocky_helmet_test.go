package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rockyHelmet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rockyHelmet, err := NewRockyHelmet()
		if err != nil {
			t.Errorf("Error creating rocky helmet: %v", err)
		}
		statBoost := rockyHelmet.GetStatBoosts()
		assert.Equal(t, stats.Stats{Hp: 270.0, Defense: 51.0}, statBoost)
	})
}
