package playable_pokemons

import (
	"pokemon-unite-damage-calculator/items"
	"pokemon-unite-damage-calculator/pokemon"
)

// Pikachu is a pokemon
type Pikachu struct {
	pokemon.Stats
	pokemon.CoolDowns
	pokemon.Buffs
}

func NewPikachu(level int, heldItems []items.HeldItemName, battleItems items.BattleItemName) (p *Pikachu, err error) {

}

func (p *Pikachu) GetAvailableAttacks() (availableAttacks []pokemon.AttackOption, err error) {
	return
}

func (p *Pikachu) Attack(attack pokemon.AttackOption) (availableAttacks []pokemon.AttackOption, err error) {
	return
}
