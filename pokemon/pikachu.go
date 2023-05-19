package pokemon

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	stats "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// Pikachu is a pokemon
type Pikachu struct {
	stats.Stats
	attack.CoolDowns
	stats.Buffs
	HeldItems []helditems.HeldItem // TODO once struct is complete
	//BattleItem []battleitems. // TODO once struct is complete
}

func NewPikachu(level int, heldItems []string, battleItem string) (p *Pikachu, err error) {
	pokemonStats, err := FetchPokemonStats(pikachuName, level)
	p.Stats = pokemonStats

	fmt.Println(pokemonStats)
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
