package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// WiseGlasses is a held item
type WiseGlasses struct {
	HeldItemData
}

// NewWiseGlasses creates a new WiseGlasses held item
func NewWiseGlasses(originalStat stats.Stats) (wiseGlasses *WiseGlasses, err error) {
	wiseGlassesData, err := FetchHeldItemData(WiseGlassesName)
	if err != nil {
		fmt.Println("Error fetching held item data for Wise Glasses")
		return
	}

	// Wise Glasses increases Special Attack by 7%
	wiseGlassesData.SpecialAttack += originalStat.SpecialAttack * 0.07
	wiseGlasses = &WiseGlasses{
		HeldItemData: wiseGlassesData,
	}
	return
}

func (a *WiseGlasses) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *WiseGlasses) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
