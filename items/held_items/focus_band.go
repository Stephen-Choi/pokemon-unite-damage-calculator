package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// FocusBand is a held item
type FocusBand struct {
	HeldItemData
}

// NewFocusBand creates a new FocusBand held item
func NewFocusBand() (focusBand *FocusBand, err error) {
	focusBandData, err := FetchHeldItemData(FocusBandName)
	if err != nil {
		fmt.Println("Error fetching held item data for Focus Band")
		return
	}

	focusBand = &FocusBand{
		HeldItemData: focusBandData,
	}
	return
}

func (a *FocusBand) GetName() string {
	return "focus band"
}

func (a *FocusBand) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *FocusBand) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
