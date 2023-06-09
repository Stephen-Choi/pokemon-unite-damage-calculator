package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// ChoiceSpecs is a held item
type ChoiceSpecs struct {
	HeldItemData
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the item is on cooldown
	used     bool    // used is a boolean which is used to check if the item has ever been used
}

// NewChoiceSpecs creates a new ChoiceSpecs held item
func NewChoiceSpecs() (choiceSpecs *ChoiceSpecs, err error) {
	choiceSpecsData, err := FetchHeldItemData(ChoiceSpecsName)
	if err != nil {
		fmt.Println("Error fetching held item data for Choice Specs")
		return
	}

	choiceSpecs = &ChoiceSpecs{
		HeldItemData: choiceSpecsData,
	}
	return
}

func (item *ChoiceSpecs) GetName() string {
	return "choice specs"
}

func (item *ChoiceSpecs) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *ChoiceSpecs) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Skip if item activation is on cooldown
	if item.isOnCooldown(elapsedTime) {
		onCooldown = true
		return
	}

	// Choice specs only activates on move1 or move2. Move must also inflict damage.
	if attackOption != attack.Move1 && attackOption != attack.Move2 && attackDamage != 0.0 {
		return // early return, don't trigger cooldown
	}

	// Perform choice specs effect
	extraDamage := 60.0 + 0.4*float64(originalStats.SpecialAttack)
	effect.AdditionalDamage = attack.AdditionalDamage{
		Type:   attack.SingleInstance,
		Amount: extraDamage,
	}

	// Put the held item on cooldown
	item.setLastUsed(elapsedTime)

	return
}

// isOnCooldown checks if the battle item is on cooldown
func (item *ChoiceSpecs) isOnCooldown(elapsedTime float64) bool {
	if !item.used {
		return false
	}
	itemCooldown := item.SpecialEffect.AdditionalDamage.Cooldown
	return item.lastUsed+itemCooldown > elapsedTime
}

// setLastUsed sets the lastUsed time to now
func (item *ChoiceSpecs) setLastUsed(elapsedTime float64) {
	item.lastUsed = elapsedTime
	item.used = true
}
