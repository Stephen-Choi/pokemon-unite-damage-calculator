package pokemon

import (
	"fmt"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon/pikachu"
)

func GetPokemon(pokemonName string, level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem) (pokemon Pokemon, err error) {
	if !IsPokemonPlayable(pokemonName) {
		err = fmt.Errorf("pokemon %s is not playable", pokemonName)
		return
	}

	switch pokemonName {
	case PikachuName:
		pokemon, err = pikachu.NewPikachu(level, move1Name, move2Name, heldItems, battleItem, nil)
	}

	if err != nil {
		err = fmt.Errorf("error getting pokemon %s: %w", pokemonName, err)
	}

	return
}
