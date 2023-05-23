package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// RescueHood is a held item
type RescueHood struct {
	HeldItemData
}

// NewRescueHood creates a new RescueHood held item
func NewRescueHood() (rescueHood *RescueHood, err error) {
	rescueHoodData, err := FetchHeldItemData(RescueHoodName)
	if err != nil {
		fmt.Println("Error fetching held item data for Rescue Hood")
		return
	}

	rescueHood = &RescueHood{
		HeldItemData: rescueHoodData,
	}
	return
}

func (a *RescueHood) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *RescueHood) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
