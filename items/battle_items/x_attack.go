package battleitems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
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

func (item *XAttack) GetName() string {
	return "x attack"
}

// Activate activates the battle item
func (item *XAttack) Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, battleItemEffect BattleItemEffect, err error) {
	// Skip if item activation is on cooldown
	if !item.IsAvailable(elapsedTime) {
		onCooldown = true
		return
	}

	// Apply stat buffs
	xAttackStatsBuff := item.SpecialEffect.StatsBuff
	buff := stats.Buff{
		StatIncrease: stats.Stats{
			Attack:        xAttackStatsBuff.AttackBuff,
			SpecialAttack: xAttackStatsBuff.SpecialAttackBuff,
			AttackSpeed:   xAttackStatsBuff.AttackSpeedBuff,
		},
		BuffType:    stats.PercentIncrease,
		DurationEnd: elapsedTime + time.ConvertSecondsToMilliseconds(xAttackStatsBuff.Duration),
	}

	battleItemEffect = BattleItemEffect{
		Buff: buff,
	}

	// Put the battle item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// IsAvailable checks if the battle item is available
func (item *XAttack) IsAvailable(elapsedTime float64) bool {
	if !item.used {
		return true
	}

	return item.lastUsed+item.Cooldown <= elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *XAttack) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
