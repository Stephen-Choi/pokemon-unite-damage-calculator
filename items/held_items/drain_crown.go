package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// DrainCrown is a held item
type DrainCrown struct {
	HeldItemData
}

// NewDrainCrown creates a new DrainCrown held item
func NewDrainCrown() (drainCrown *DrainCrown, err error) {
	drainCrownData, err := FetchHeldItemData(DrainCrownName)
	if err != nil {
		fmt.Println("Error fetching held item data for Drain Crown")
		return
	}

	drainCrown = &DrainCrown{
		HeldItemData: drainCrownData,
	}
	return
}

func (a *DrainCrown) GetName() string {
	return "drain crown"
}

func (a *DrainCrown) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *DrainCrown) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
