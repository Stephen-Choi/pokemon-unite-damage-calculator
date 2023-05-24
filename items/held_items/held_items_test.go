package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testCooldown tests the cooldown of a battle item
func testCooldown(t *testing.T, heldItem HeldItem, itemCooldown float64, attackOption attack.Option, attackType attack.Type) {
	// Starting time
	startingTime := 0.0

	// Activate
	onCooldown, _, err := heldItem.Activate(stats.Stats{}, startingTime, attackOption, attackType)
	assert.NoError(t, err)
	assert.False(t, onCooldown)

	// Assert that item will not activate if on cooldown
	onCooldownTime := itemCooldown - 1
	onCooldown, _, err = heldItem.Activate(stats.Stats{}, onCooldownTime, attackOption, attackType)
	assert.NoError(t, err)
	assert.True(t, onCooldown)

	// Assert that item will activate after cooldown
	offCooldownTime := itemCooldown
	onCooldown, _, err = heldItem.Activate(stats.Stats{}, offCooldownTime, attackOption, attackType)
	assert.NoError(t, err)
	assert.False(t, onCooldown)
}

// Test_cooldowns tests the cooldowns of all battle items
func Test_cooldowns(t *testing.T) {
	t.Run("ChoiceSpecs", func(t *testing.T) {
		choiceSpecs, err := NewChoiceSpecs()
		assert.NoError(t, err)
		testCooldown(t, choiceSpecs, choiceSpecs.SpecialEffect.AdditionalDamage.Cooldown, attack.Move1, attack.PhysicalAttack)
	})
	t.Run("Energy Amplifier", func(t *testing.T) {
		energyAmplifier, err := NewEnergyAmplifier()
		assert.NoError(t, err)
		testCooldown(t, energyAmplifier, energyAmplifier.SpecialEffect.Buff.Cooldown, attack.UniteMove, attack.PhysicalAttack)
	})
	t.Run("Muscle Band", func(t *testing.T) {
		muscleBand, err := NewMuscleBand()
		assert.NoError(t, err)
		testCooldown(t, muscleBand, muscleBand.SpecialEffect.AdditionalDamage.InternalCooldown, attack.BasicAttack, attack.PhysicalAttack)
	})
	t.Run("Razor Claw", func(t *testing.T) {
		razorClaw, err := NewRazorClaw()
		assert.NoError(t, err)
		testCooldown(t, razorClaw, razorClaw.SpecialEffect.AdditionalDamage.Cooldown, attack.Move2, attack.PhysicalAttack)
	})
	t.Run("Scope Lens", func(t *testing.T) {
		scopeLens, err := NewScopeLens()
		assert.NoError(t, err)
		testCooldown(t, scopeLens, scopeLens.SpecialEffect.AdditionalDamage.Cooldown, attack.CriticalHitBasicAttack, attack.PhysicalAttack)
	})
}
