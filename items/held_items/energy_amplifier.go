package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
	"github.com/samber/lo"
)

// EnergyAmplifier is a held item
type EnergyAmplifier struct {
	HeldItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewEnergyAmplifier creates a new EnergyAmplifier held item
func NewEnergyAmplifier() (energyAmplifier *EnergyAmplifier, err error) {
	energyAmplifierData, err := FetchHeldItemData(EnergyAmplifierName)
	if err != nil {
		fmt.Println("Error fetching held item data for Energy Amplifier")
		return
	}

	energyAmplifier = &EnergyAmplifier{
		HeldItemData: energyAmplifierData,
	}
	return
}

func (item *EnergyAmplifier) GetName() string {
	return "energy amplifier"
}

func (item *EnergyAmplifier) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *EnergyAmplifier) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Energy amplifier only activates on unite move
	if attackOption != attack.UniteMove {
		return // early return, don't trigger cooldown
	}

	fmt.Println("EnergyAmplifier.Activate: activated")

	// Perform energy amplifier effect
	effect.AdditionalDamage = attack.AdditionalDamage{
		Type:        attack.PercentDamageBoost,
		Amount:      0.21,
		DurationEnd: lo.ToPtr(time.ConvertSecondsToMilliseconds(item.SpecialEffect.Buff.Duration) + elapsedTime),
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *EnergyAmplifier) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}
	itemCooldown := item.SpecialEffect.Buff.Cooldown
	return item.lastUsed+itemCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *EnergyAmplifier) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
