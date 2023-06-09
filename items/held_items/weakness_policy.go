package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
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

func (a *WeaknessPolicy) GetName() string {
	return "weakness policy"
}

func (a *WeaknessPolicy) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *WeaknessPolicy) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
