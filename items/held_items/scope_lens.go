package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// ScopeLens is a held item
type ScopeLens struct {
	HeldItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewScopeLens creates a new ScopeLens held item
func NewScopeLens() (scopeLens *ScopeLens, err error) {
	scopeLensData, err := FetchHeldItemData(ScopeLensName)
	if err != nil {
		fmt.Println("Error fetching held item data for Scope Lens")
		return
	}

	scopeLens = &ScopeLens{
		HeldItemData: scopeLensData,
	}
	return
}

func (item *ScopeLens) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *ScopeLens) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.AttackOption) (onCooldown bool, effect HeldItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Scope Lens only activates on critical hit basic attack
	if attackOption != attack.CriticalHitBasicAttack {
		return // early return, don't trigger cooldown
	}

	// Perform Scope Lens effect
	extraDamage := 0.75 * float64(originalStats.Attack)
	effect.AdditionalDamage = attack.AdditionalDamage{
		Type:   attack.SimpleAdditionalDamage,
		Amount: extraDamage,
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *ScopeLens) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}
	itemCooldown := item.SpecialEffect.AdditionalDamage.Cooldown
	return item.lastUsed+itemCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *ScopeLens) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
