package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rustedSword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rustedSword, err := NewRustedSword()
		if err != nil {
			t.Errorf("Error creating rusted sword: %v", err)
		}
		statBoost := rustedSword.GetStatBoosts()
		assert.Equal(t, stats.Stats{}, statBoost)
	})
}
