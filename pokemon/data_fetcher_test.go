package pokemon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_fetchPokemonStats tests the FetchPokemonStats function
func Test_fetchPokemonStats(t *testing.T) {
	for _, playablePokemon := range PlayablePokemons {
		for level := 1; level < maxLevel; level++ {
			_, err := FetchPokemonStats(playablePokemon, level)
			assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at level : %d", playablePokemon, level))
		}
	}
}
