package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
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

func (a *RockyHelmet) GetName() string {
	return "rocky helmet"
}

func (a *RockyHelmet) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *RockyHelmet) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
