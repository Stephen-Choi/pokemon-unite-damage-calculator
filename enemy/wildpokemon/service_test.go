package wildpokemon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetObjectivePokemon(t *testing.T) {
	t.Run("success retrieving all objective pokemon at all remaining times", func(t *testing.T) {
		for _, objectivePokemon := range ObjectivePokemonNames {
			if objectivePokemon == RayquazaName {
				for _, remainingTime := range validRayquazaRemainingTimes {
					_, err := GetObjectivePokemon(objectivePokemon, remainingTime)
					assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at remaining time: %d", objectivePokemon, remainingTime))
				}
			} else if objectivePokemon == RegielekiNeutralName || objectivePokemon == RegielekiEnemyName || objectivePokemon == BottomRegis {
				for _, remainingTime := range validRegiRemainingTimes {
					_, err := GetObjectivePokemon(objectivePokemon, remainingTime)
					assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at remaining time: %d", objectivePokemon, remainingTime))
				}
			}
		}
	})
}
