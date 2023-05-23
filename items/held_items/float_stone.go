package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// FloatStone is a held item
type FloatStone struct {
	HeldItemData
}

// NewFloatStone creates a new FloatStone held item
func NewFloatStone() (floatStone *FloatStone, err error) {
	floatStoneData, err := FetchHeldItemData(FloatStoneName)
	if err != nil {
		fmt.Println("Error fetching held item data for Float Stone")
		return
	}

	floatStone = &FloatStone{
		HeldItemData: floatStoneData,
	}
	return
}

func (a *FloatStone) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *FloatStone) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
