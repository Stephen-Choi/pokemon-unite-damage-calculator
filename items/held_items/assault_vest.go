package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// AssaultVest is a held item
type AssaultVest struct {
	HeldItemData
}

// NewAssaultVest creates a new AssaultVest held item
func NewAssaultVest() (assaultVest *AssaultVest, err error) {
	assaultVestData, err := FetchHeldItemData(AssaultVestName)
	if err != nil {
		fmt.Println("Error fetching held item data for Assault Vest")
		return
	}

	assaultVest = &AssaultVest{
		HeldItemData: assaultVestData,
	}
	return
}

func (a *AssaultVest) GetName() string {
	return "assault vest"
}

func (a *AssaultVest) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *AssaultVest) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
