package fetchdata

import (
	"encoding/json"
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"io/ioutil"
	"path/filepath"
)

const (
	pokemonStatsDirectory = "pokemon_stats"
)

// fetchPokemonStats fetches the stats of a pokemon at a given level
func fetchPokemonStats(pokemonName string, level int) (stats stats.Stats, err error) {
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
