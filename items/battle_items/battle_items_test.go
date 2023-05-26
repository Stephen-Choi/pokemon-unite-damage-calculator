package battleitems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testCooldown tests the cooldown of a battle item
func testCooldown(t *testing.T, battleItem BattleItem, itemCooldown float64) {
	// Starting time
	startingTime := 0.0

	// Activate
	onCooldown, _, err := battleItem.Activate(stats.Stats{}, startingTime)
	assert.NoError(t, err)
	assert.False(t, onCooldown)

	// Assert that item will not activate if on cooldown
	onCooldownTime := itemCooldown - 1
	onCooldown, _, err = battleItem.Activate(stats.Stats{}, onCooldownTime)
	assert.NoError(t, err)
	assert.True(t, onCooldown)

	// Assert that item will activate after cooldown
	offCooldownTime := itemCooldown
	onCooldown, _, err = battleItem.Activate(stats.Stats{}, offCooldownTime)
	assert.NoError(t, err)
	assert.False(t, onCooldown)
}

// Test_cooldowns tests the time of all battle items
func Test_cooldowns(t *testing.T) {
	t.Run("FluffyTail", func(t *testing.T) {
		fluffyTail, err := NewFluffyTail()
		assert.NoError(t, err)
		testCooldown(t, fluffyTail, fluffyTail.Cooldown)
	})
	t.Run("XAttack", func(t *testing.T) {
		xAttack, err := NewXAttack()
		assert.NoError(t, err)
		testCooldown(t, xAttack, xAttack.Cooldown)
	})
}
