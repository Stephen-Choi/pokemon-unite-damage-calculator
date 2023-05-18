package pokemon

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	stats2 "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// Pikachu is a pokemon
type Pikachu struct {
	stats2.Stats
	attack.CoolDowns
	stats2.Buffs
	HeldItems  // TODO once struct is complete
	BattleItem // TODO once struct is complete
}

func NewPikachu(level int, heldItems []helditems.Name, battleItem battleitems.Name) (p *Pikachu, err error) {
	stats, err := stats2.fetchPokemonStats(pikachuName, level)
	p.Stats = stats

	fmt.Println(stats)
	fmt.Println(heldItems)
	fmt.Println(battleItem)

	return
}

func (p *Pikachu) GetAvailableAttacks() (availableAttacks []attack.AttackOption, err error) {
	return
}

func (p *Pikachu) Attack(attack attack.AttackOption) (availableAttacks []attack.AttackOption, err error) {
	return
}
