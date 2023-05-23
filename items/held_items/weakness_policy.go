package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// WeaknessPolicy is a held item
type WeaknessPolicy struct {
	HeldItemData
}

// NewWeaknessPolicy creates a new WeaknessPolicy held item
func NewWeaknessPolicy() (weaknessPolicy *WeaknessPolicy, err error) {
	weaknessPolicyData, err := FetchHeldItemData(WeaknessPolicyName)
	if err != nil {
		fmt.Println("Error fetching held item data for Weakness Policy")
		return
	}

	weaknessPolicy = &WeaknessPolicy{
		HeldItemData: weaknessPolicyData,
	}
	return
}

func (a *WeaknessPolicy) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *WeaknessPolicy) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
