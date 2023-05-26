package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// RapidFireScarf is a held item
type RapidFireScarf struct {
	HeldItemData
	internalStackCounter int
	lastStackAppliedTime float64 // lastStackAppliedTime is a time in milliseconds which is used to check when the last stack was applied
	lastUsed             float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used                 bool    // used is a boolean which is used to check if the item has ever been used
}

// NewRapidFireScarf creates a new RapidFireScarf held item
func NewRapidFireScarf() (scarf *RapidFireScarf, err error) {
	scarfData, err := FetchHeldItemData(RapidFireScarfName)
	if err != nil {
		fmt.Println("Error fetching held item data for Rapid Fire Scarf")
		return
	}

	scarf = &RapidFireScarf{
		HeldItemData: scarfData,
	}
	return
}

func (item *RapidFireScarf) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *RapidFireScarf) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Rapid Fire Scarf only holds internal stacks for 3s
	if elapsedTime > item.lastStackAppliedTime+3000 {
		item.internalStackCounter = 0
	}

	// Rapid Fire Scarf only activates on basic attack
	if attackOption != attack.BasicAttackOption && attackOption != attack.CriticalHitBasicAttack {
		return // early return, don't trigger cooldown
	}

	// Rapid Fire Scarf only activates on the 3rd stack
	if item.internalStackCounter < 3 {
		item.internalStackCounter++
		item.lastStackAppliedTime = elapsedTime
		return
	}

	// Perform Rapid Fire Scarf effect
	effect.UpdatedStats = stats.StatBuff{
		Stats: stats.Stats{
			AttackSpeed: 0.3 * float64(originalStats.AttackSpeed),
		},
		DurationEnd: elapsedTime + item.SpecialEffect.Buff.Duration,
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *RapidFireScarf) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}
	itemCooldown := item.SpecialEffect.Buff.Cooldown
	itemInternalCooldown := item.SpecialEffect.Buff.InternalCooldown
	return item.lastUsed+itemCooldown > elapsedTime || item.lastStackAppliedTime+itemInternalCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *RapidFireScarf) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
