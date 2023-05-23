package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// RockyHelmet is a held item
type RockyHelmet struct {
	HeldItemData
}

// NewRockyHelmet creates a new RockyHelmet held item
func NewRockyHelmet() (rockyHelmet *RockyHelmet, err error) {
	rockyHelmetData, err := FetchHeldItemData(RockyHelmetName)
	if err != nil {
		fmt.Println("Error fetching held item data for Rocky Helmet")
		return
	}

	rockyHelmet = &RockyHelmet{
		HeldItemData: rockyHelmetData,
	}
	return
}

func (a *RockyHelmet) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *RockyHelmet) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
