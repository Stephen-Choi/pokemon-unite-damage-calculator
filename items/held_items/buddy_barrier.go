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

func (a *BuddyBarrier) GetName() string {
	return "buddy barrier"
}

func (a *BuddyBarrier) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	return a.Stats
}

func (a *BuddyBarrier) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
