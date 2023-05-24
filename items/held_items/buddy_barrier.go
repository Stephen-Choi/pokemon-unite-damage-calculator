package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// BuddyBarrier is a held item
type BuddyBarrier struct {
	HeldItemData
}

// NewBuddyBarrier creates a new BuddyBarrier held item
func NewBuddyBarrier() (buddyBarrier *BuddyBarrier, err error) {
	buddyBarrierData, err := FetchHeldItemData(BuddyBarrierName)
	if err != nil {
		fmt.Println("Error fetching held item data for Buddy Barrier")
		return
	}

	buddyBarrier = &BuddyBarrier{
		HeldItemData: buddyBarrierData,
	}
	return
}

func (a *BuddyBarrier) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *BuddyBarrier) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
