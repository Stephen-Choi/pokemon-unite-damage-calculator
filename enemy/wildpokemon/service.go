package wildpokemon

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
)

// GetObjectivePokemon fetches an enemy wild pokemon
func GetObjectivePokemon(pokemonName string, remainingTime int) (enemyWildPokemon enemy.Pokemon, err error) {
	err = IsValidObjectivePokemon(pokemonName, remainingTime)
	if err != nil {
		return
	}

	switch pokemonName {
	case RayquazaName:
		enemyWildPokemon, err = NewRayquaza(remainingTime)
	case RegielekiNeutralName:
		enemyWildPokemon, err = NewRegielekiNeutral(remainingTime)
	case RegielekiEnemyName:
		enemyWildPokemon, err = NewRegielekiEnemy(remainingTime)
	case BottomRegis:
		enemyWildPokemon, err = NewBottomRegi(remainingTime)
	}
	return
}
