package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// Leftovers is a held item
type Leftovers struct {
	HeldItemData
}

// NewLeftovers creates a new Leftovers held item
func NewLeftovers() (leftovers *Leftovers, err error) {
	leftoversData, err := FetchHeldItemData(LeftoversName)
	if err != nil {
		fmt.Println("Error fetching held item data for Leftovers")
		return
	}

	leftovers = &Leftovers{
		HeldItemData: leftoversData,
	}
	return
}

func (a *Leftovers) GetName() string {
	return "leftovers"
}

func (a *Leftovers) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *Leftovers) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
