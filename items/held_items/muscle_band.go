package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/samber/lo"
)

// MuscleBand is a held item
type MuscleBand struct {
	HeldItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewMuscleBand creates a new MuscleBand held item
func NewMuscleBand() (muscleBand *MuscleBand, err error) {
	muscleBandData, err := FetchHeldItemData(MuscleBandName)
	if err != nil {
		fmt.Println("Error fetching held item data for Muscle Band")
		return
	}

	muscleBand = &MuscleBand{
		HeldItemData: muscleBandData,
	}
	return
}

func (item *MuscleBand) GetName() string {
	return "muscle band"
}

func (item *MuscleBand) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *MuscleBand) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Muscle Band only activates on basic attack
	if attackOption != attack.BasicAttackOption && attackOption != attack.CriticalHitBasicAttack {
		return // early return, don't trigger cooldown
	}

	// Perform Muscle Band effect
	effect.AdditionalDamage = attack.AdditionalDamage{
		Type:         attack.RemainingEnemyHp,
		Amount:       0.03,
		CappedAmount: lo.ToPtr(360.0),
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *MuscleBand) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}
	itemCooldown := item.SpecialEffect.AdditionalDamage.InternalCooldown
	return item.lastUsed+itemCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *MuscleBand) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
