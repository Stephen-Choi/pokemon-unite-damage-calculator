package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// SpecialAttackSpecs is a held item that increases health stats based on stacks
type SpecialAttackSpecs struct {
	HeldItemData
	numStacks int
}

// NewSpecialAttackSpecs creates a new SpecialAttackSpecs held item
func NewSpecialAttackSpecs(numStacks int) (specialAttackSpecs *SpecialAttackSpecs, err error) {
	specialAttackSpecsData, err := FetchHeldItemData(SpecialAttackSpecsName)
	if err != nil {
		fmt.Println("Error fetching held item data for Special Attack Specs")
		return
	}

	specialAttackSpecs = &SpecialAttackSpecs{
		HeldItemData: specialAttackSpecsData,
		numStacks:    numStacks,
	}
	return
}

func (a *SpecialAttackSpecs) GetStatBoosts() stats.Stats {
	stackBoosts := a.numStacks * a.SpecialEffect.Stack.Amount
	a.Stats.SpecialAttack += float64(stackBoosts)
	return a.Stats
}

func (a *SpecialAttackSpecs) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
