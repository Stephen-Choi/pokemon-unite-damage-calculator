package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buddyBarrier(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		buddyBarrier, err := NewBuddyBarrier()
		if err != nil {
			t.Errorf("Error creating buddy barrier: %v", err)
		}
		statBoost := buddyBarrier.GetStatBoosts(stats.Stats{})
		assert.Equal(t, stats.Stats{Hp: 450}, statBoost)
	})
}
