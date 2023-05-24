package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rescueHood(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rescueHood, err := NewRescueHood()
		if err != nil {
			t.Errorf("Error creating rescue hood: %v", err)
		}
		statBoost := rescueHood.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Defense: 30.0, SpecialDefense: 30.0}, statBoost)
	})
}
