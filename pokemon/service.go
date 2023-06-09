package pokemon

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon/pikachu"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon/slowbro"
)

func GetPokemon(pokemonName string, level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem) (pokemon Pokemon, err error) {
	if !isPokemonPlayable(pokemonName) {
		err = fmt.Errorf("pokemon %s is not playable", pokemonName)
		return
	}

	switch pokemonName {
	case PikachuName:
		pokemon, err = pikachu.NewPikachu(level, move1Name, move2Name, heldItems, battleItem, nil)
	case SlowbroName:
		pokemon, err = slowbro.NewSlowbro(level, move1Name, move2Name, heldItems, battleItem, nil)
	}

	if isPokemonIsDefenderOrSupporter(pokemonName) {
		pokemon.AddAdditionalDamage("defender-and-supporter-wild-pokemon-damage-boost", attack.GetDefenderAndSupporterWildDamageBoost(level))
	}

	return
}
