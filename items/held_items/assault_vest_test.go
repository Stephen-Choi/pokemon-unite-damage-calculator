package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_assaultVest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assaultVest, err := NewAssaultVest()
		if err != nil {
			t.Errorf("Error creating assault vest: %v", err)
		}
		statBoost := assaultVest.GetStatBoosts()
		assert.Equal(t, stats.Stats{Hp: 270, SpecialDefense: 51}, statBoost)
	})
}
