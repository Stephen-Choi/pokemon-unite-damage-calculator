package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// WiseGlasses is a held item
type WiseGlasses struct {
	HeldItemData
}

// NewWiseGlasses creates a new WiseGlasses held item
func NewWiseGlasses() (wiseGlasses *WiseGlasses, err error) {
	wiseGlassesData, err := FetchHeldItemData(WiseGlassesName)
	if err != nil {
		fmt.Println("Error fetching held item data for Wise Glasses")
		return
	}

	wiseGlasses = &WiseGlasses{
		HeldItemData: wiseGlassesData,
	}

	return
}

func (a *WiseGlasses) GetName() string {
	return "wise glasses"
}

func (a *WiseGlasses) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	// Wise Glasses increases Special Attack by 7%
	a.SpecialAttack += originalStats.SpecialAttack * 0.07
	return a.Stats
}

func (a *WiseGlasses) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
