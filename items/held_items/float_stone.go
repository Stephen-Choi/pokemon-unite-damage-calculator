package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
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

func (a *FloatStone) GetName() string {
	return "float stone"
}

func (a *FloatStone) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *FloatStone) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
