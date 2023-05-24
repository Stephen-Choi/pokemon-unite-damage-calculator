package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// AttackWeight is a held item that increases attack stats based on stacks
type AttackWeight struct {
	HeldItemData
	numStacks int
}

// NewAttackWeight creates a new AttackWeight held item
func NewAttackWeight(numStacks int) (attackWeight *AttackWeight, err error) {
	attackWeightData, err := FetchHeldItemData(AttackWeightName)
	if err != nil {
		fmt.Println("Error fetching held item data for Attack Weight")
		return
	}

	attackWeight = &AttackWeight{
		HeldItemData: attackWeightData,
		numStacks:    numStacks,
	}
	return
}

func (a *AttackWeight) GetStatBoosts() stats.Stats {
	stackBoosts := a.numStacks * a.SpecialEffect.Stack.Amount
	a.Stats.Attack += float64(stackBoosts)
	return a.Stats
}

func (a *AttackWeight) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
