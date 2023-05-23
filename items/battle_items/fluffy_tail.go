package battleitems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// FluffyTail is a battle item that deals additional damage to a target for a short period of time
type FluffyTail struct {
	BattleItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewFluffyTail creates a new FluffyTail battle item
func NewFluffyTail() (fluffyTail *FluffyTail, err error) {
	fluffyTailData, err := FetchBattleItemData(FluffyTailName)
	if err != nil {
		fmt.Println("Error fetching battle item data for Fluffy Tail")
		return
	}

	fluffyTail = &FluffyTail{
		BattleItemData: fluffyTailData,
	}
	return
}

// Activate activates the battle item
func (item *FluffyTail) Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, effect BattleItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Additional damage formula: 100% attack + 60% sp. attack + 10(level-1) + 100"
	extraDamage := 1.0*originalStats.Attack + 0.6*originalStats.SpecialAttack + 10.0*(float64(originalStats.Level-1)) + 100.0
	effect.AdditionalDamage.Damage = extraDamage
	effect.AdditionalDamage.Duration = item.SpecialEffect.AdditionalDamage.Duration

	// Put the battle item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *FluffyTail) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}

	return item.lastUsed+item.Cooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *FluffyTail) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
