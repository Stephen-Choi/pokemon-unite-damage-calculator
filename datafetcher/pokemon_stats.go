package datafetcher

import (
	"encoding/json"
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"io/ioutil"
	"path/filepath"
)

const (
	pokemonStatsDirectory = "/Users/stephenchoi/Desktop/Projects/pokemon-unite-damage-calculator/data/pokemon_stats"
)

// FetchPokemonStats fetches the stats of a pokemon at a given level
func FetchPokemonStats(pokemonName string, level int) (pokemonStats stats.Stats, err error) {
	pokemonFileName := pokemonName + "_stats.json"

	// Read file
	pokemonFilePath := filepath.Join(pokemonStatsDirectory, pokemonFileName)
	data, err := ioutil.ReadFile(pokemonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Pokemon JsonStats
	var pokemonJsonStats []stats.JsonStats
	err = json.Unmarshal(data, &pokemonJsonStats)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Convert the json stats for the specified level to internal use struct (with better typing)
	index := level - 1 // level 1 is at index 0
	return stats.ToTypedStats(pokemonJsonStats[index])
}
