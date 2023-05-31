package battleitems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
	"github.com/samber/lo"
)

// FluffyTail is a battle item that deals additional damage to a target for a short period of time
type FluffyTail struct {
	BattleItemData
	internalCooldown float64
	lastUsed         float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used             bool    // used is a boolean which is used to check if the item has ever been used
}

// NewFluffyTail creates a new FluffyTail battle item
func NewFluffyTail() (fluffyTail *FluffyTail, err error) {
	fluffyTailData, err := FetchBattleItemData(FluffyTailName)
	if err != nil {
		fmt.Println("Error fetching battle item data for Fluffy Tail")
		return
	}

	fluffyTail = &FluffyTail{
		BattleItemData:   fluffyTailData,
		internalCooldown: 500,
	}
	return
}

// Activate activates the battle item
func (item *FluffyTail) Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, battleItemEffect BattleItemEffect, err error) {
	// Skip if item activation is on cooldown
	if !item.IsAvailable(elapsedTime) {
		onCooldown = true
		return
	}

	// Additional damage formula: 100% attack + 60% sp. attack + 10(level-1) + 100"
	damage := 1.0*originalStats.Attack + 0.6*originalStats.SpecialAttack + 10.0*(float64(originalStats.Level-1)) + 100.0
	additionalDamage := attack.AdditionalDamage{
		Type:        attack.SimpleAdditionalDamage,
		Amount:      damage,
		DurationEnd: lo.ToPtr(time.ConvertSecondsToMilliseconds(item.SpecialEffect.AdditionalDamage.Duration) + elapsedTime),
	}

	battleItemEffect = BattleItemEffect{
		AdditionalDamage: additionalDamage,
	}

	// Put the battle item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// IsAvailable checks if the battle item is available
func (item *FluffyTail) IsAvailable(elapsedTime float64) bool {
	if !item.used {
		return true
	}

	return item.lastUsed+item.Cooldown <= elapsedTime && item.lastUsed+item.internalCooldown <= elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *FluffyTail) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
