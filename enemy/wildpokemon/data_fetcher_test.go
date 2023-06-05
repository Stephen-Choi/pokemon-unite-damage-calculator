package wildpokemon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_fetchWildPokemonData tests the fetchWildPokemonData function
func Test_fetchWildPokemonData(t *testing.T) {
	for _, objectivePokemon := range ObjectivePokemonNames {
		if objectivePokemon == RayquazaName {
			for _, remainingTime := range validRayquazaRemainingTimes {
				_, err := fetchWildPokemonData(objectivePokemon, remainingTime)
				assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at remaining time: %d", objectivePokemon, remainingTime))
			}
		} else if objectivePokemon == RegielekiNeutralName || objectivePokemon == RegielekiEnemyName || objectivePokemon == BottomRegis {
			for _, remainingTime := range validRegiRemainingTimes {
				_, err := fetchWildPokemonData(objectivePokemon, remainingTime)
				assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at remaining time: %d", objectivePokemon, remainingTime))
			}
		}
	}
}
