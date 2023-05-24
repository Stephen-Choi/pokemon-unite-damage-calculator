package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testCooldown tests the cooldown of a battle item
func testCooldown(t *testing.T, heldItem HeldItem, itemCooldown float64, attackOption attack.AttackOption) {
	// Starting time
	startingTime := 0.0

	// Activate
	onCooldown, _, err := heldItem.Activate(stats.Stats{}, startingTime, attackOption)
	assert.NoError(t, err)
	assert.False(t, onCooldown)

	// Assert that item will not activate if on cooldown
	onCooldownTime := itemCooldown - 1
	onCooldown, _, err = heldItem.Activate(stats.Stats{}, onCooldownTime, attackOption)
	assert.NoError(t, err)
	assert.True(t, onCooldown)

	// Assert that item will activate after cooldown
	offCooldownTime := itemCooldown
	onCooldown, _, err = heldItem.Activate(stats.Stats{}, offCooldownTime, attackOption)
	assert.NoError(t, err)
	assert.False(t, onCooldown)
}

// Test_cooldowns tests the cooldowns of all battle items
func Test_cooldowns(t *testing.T) {
	t.Run("ChoiceSpecs", func(t *testing.T) {
		choiceSpecs, err := NewChoiceSpecs()
		assert.NoError(t, err)
		testCooldown(t, choiceSpecs, choiceSpecs.SpecialEffect.AdditionalDamage.Cooldown, attack.Move1)
	})
	t.Run("Energy Amplifier", func(t *testing.T) {
		energyAmplifier, err := NewEnergyAmplifier()
		assert.NoError(t, err)
		testCooldown(t, energyAmplifier, energyAmplifier.SpecialEffect.Buff.Cooldown, attack.UniteMove)
	})
}