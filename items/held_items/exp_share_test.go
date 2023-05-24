package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_expShare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expShare, err := NewExpShare()
		if err != nil {
			t.Errorf("Error creating score shield: %v", err)
		}
		statBoost := expShare.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Hp: 240}, statBoost)
	})
}
