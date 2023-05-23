package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// ShellBell is a held item
type ShellBell struct {
	HeldItemData
}

// NewShellBell creates a new ShellBell held item
func NewShellBell() (shellBell *ShellBell, err error) {
	shellBellData, err := FetchHeldItemData(ShellBellName)
	if err != nil {
		fmt.Println("Error fetching held item data for Shell Bell")
		return
	}

	shellBell = &ShellBell{
		HeldItemData: shellBellData,
	}
	return
}

func (a *ShellBell) GetStatBoosts() stats.Stats {
	return a.Stats
}

func (a *ShellBell) Activate(originalStats stats.Stats) (onCooldown bool, effect HeldItemEffect, err error) {
	panic("Not implemented")
}
