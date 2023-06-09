package helditems

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// SlickSpoon is a held item
type SlickSpoon struct {
	HeldItemData
}

// NewSlickSpoon creates a new SlickSpoon held item
func NewSlickSpoon() (slickSpoon *SlickSpoon, err error) {
	slickSpoonData, err := FetchHeldItemData(SlickSpoonName)
	if err != nil {
		fmt.Println("Error fetching held item data for Slick Spoon")
		return
	}

	slickSpoon = &SlickSpoon{
		HeldItemData: slickSpoonData,
	}
	return
}

func (item *SlickSpoon) GetName() string {
	return "slick spoon"
}

func (item *SlickSpoon) GetStatBoosts(originalStats stats.Stats) stats.Stats {
	return item.Stats
}

func (item *SlickSpoon) Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error) {
	// Slick Spoon only activates on special attack type attacks
	if attackType != attack.SpecialAttack {
		return // early return
	}

	effect.Debuff = attack.Debuff{
		DebuffEffect: attack.IgnoreDefenseForAttackingPokemon,
		Stats:        stats.Stats{SpecialDefense: 0.2},
		DebuffType:   attack.Percent,
	}

	return
}
