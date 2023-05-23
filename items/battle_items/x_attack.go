package battleitems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

type XAttack struct {
	BattleItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewXAttack creates a new XAttack battle item
func NewXAttack() (xAttack *XAttack, err error) {
	xAttackData, err := FetchBattleItemData(XAttackName)
	if err != nil {
		fmt.Println("Error fetching battle item data for XAttack")
		return
	}

	xAttack = &XAttack{
		BattleItemData: xAttackData,
	}
	return
}

// Activate activates the battle item
func (item *XAttack) Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, effect BattleItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Apply stat buffs
	updatedStats := originalStats
	xAttackStatsBuff := item.SpecialEffect.StatsBuff
	updatedStats.Attack *= 1.0 + xAttackStatsBuff.AttackBuff
	updatedStats.SpecialAttack *= 1.0 + xAttackStatsBuff.SpecialAttackBuff
	updatedStats.AttackSpeed += xAttackStatsBuff.AttackSpeedBuff
	effect.UpdatedStats = updatedStats

	// Put the battle item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *XAttack) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}

	return item.lastUsed+item.Cooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *XAttack) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
