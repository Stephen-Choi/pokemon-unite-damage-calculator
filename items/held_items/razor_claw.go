package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
	"github.com/samber/lo"
)

// RazorClaw is a held item
type RazorClaw struct {
	HeldItemData
	active   bool    // active is a boolean which is used to check if the item is currently active
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewRazorClaw creates a new RazorClaw held item
func NewRazorClaw() (razorClaw *RazorClaw, err error) {
	razorClawData, err := FetchHeldItemData(RazorClawName)
	if err != nil {
		fmt.Println("Error fetching held item data for Razor Claw")
		return
	}

	razorClaw = &RazorClaw{
		HeldItemData: razorClawData,
	}
	return
}

func (item *RazorClaw) GetName() string {
	return "razor claw"
}

func (item *RazorClaw) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *RazorClaw) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	if !item.isActive(elapsedTime) {
		item.active = false
	}

	// Skip if item activation is on cooldown and item is not currently active
	if item.isOnCooldown(elapsedTime) && !item.active {
		onCooldown = true
		return
	}

	// Razor Claw only activates on move1 or move2
	if attackOption != attack.Move1 && attackOption != attack.Move2 {
		return // early return, don't trigger cooldown
	}

	// Perform Razor Claw effect
	extraDamage := 20.0 + 0.5*float64(originalStats.Attack)
	effect.AdditionalDamage = attack.AdditionalDamage{
		Type:        attack.SingleInstance,
		Amount:      extraDamage,
		DurationEnd: lo.ToPtr(time.ConvertSecondsToMilliseconds(item.SpecialEffect.AdditionalDamage.Duration + elapsedTime)),
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

func (item *RazorClaw) isActive(elapsedTime float64) bool {
	if !item.active {
		return false
	}

	// If last used + item active duration is greater than current time, it means it's still in effect
	if item.lastUsed+time.ConvertSecondsToMilliseconds(item.SpecialEffect.AdditionalDamage.Duration) > elapsedTime {
		return true
	} else {
		return false
	}
}

// isOnCooldown checks if the battle item is on cooldown
func (item *RazorClaw) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}

	itemCooldown := item.SpecialEffect.AdditionalDamage.Cooldown
	return item.lastUsed+itemCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *RazorClaw) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
