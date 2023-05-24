package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// RustedSword is a held item
type RustedSword struct {
	HeldItemData
}

// NewRustedSword creates a new RustedSword held item
func NewRustedSword() (rustedSword *RustedSword, err error) {
	rustedSwordData, err := FetchHeldItemData(RustedSwordName)
	if err != nil {
		fmt.Println("Error fetching held item data for Rusted Sword")
		return
	}

	rustedSword = &RustedSword{
		HeldItemData: rustedSwordData,
	}
	return
}

func (a *RustedSword) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *RustedSword) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
