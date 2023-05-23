package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// ExpShare is a held item
type ExpShare struct {
	HeldItemData
}

// NewExpShare creates a new ExpShare held item
func NewExpShare() (expShare *ExpShare, err error) {
	expShareData, err := FetchHeldItemData(ExpShareName)
	if err != nil {
		fmt.Println("Error fetching held item data for Exp Share")
		return
	}

	expShare = &ExpShare{
		HeldItemData: expShareData,
	}
	return
}

func (a *ExpShare) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *ExpShare) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}