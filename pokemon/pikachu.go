package pokemon

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items"
)

// Pikachu is a pokemon
type Pikachu struct {
	Stats
	CoolDowns
	Buffs
}

func NewPikachu(level int, heldItems []items.HeldItemName, battleItem items.BattleItemName) (p *Pikachu, err error) {
	stats, err := fetchPokemonStats(pikachuName, level)

	fmt.Println(stats)
	fmt.Println(heldItems)
	fmt.Println(battleItem)

	return
}

func (p *Pikachu) GetAvailableAttacks() (availableAttacks []AttackOption, err error) {
	return
}

func (p *Pikachu) Attack(attack AttackOption) (availableAttacks []AttackOption, err error) {
	return
}
