package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// AeosCookie is a held item that increases health stats based on stacks
type AeosCookie struct {
	HeldItemData
	numStacks int
}

// NewAeosCookie creates a new AeosCookie held item
func NewAeosCookie(numStacks int) (aeosCookie *AeosCookie, err error) {
	aeosCookieData, err := FetchHeldItemData(AeosCookieName)
	if err != nil {
		fmt.Println("Error fetching held item data for Aeos Cookie")
		return
	}

	aeosCookie = &AeosCookie{
		HeldItemData: aeosCookieData,
		numStacks:    numStacks,
	}
	return
}

func (a *AeosCookie) GetName() string {
	return "aeos cookie"
}

func (a *AeosCookie) GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats) {
	stackBoosts := a.numStacks * a.SpecialEffect.Stack.Amount
	a.Stats.Hp += float64(stackBoosts)
	return a.Stats
}

func (a *AeosCookie) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type, attackDamage float64) (onCooldown bool, effect HeldItemEffect, err error) {
	// Not damage related, simple return
	return
}
