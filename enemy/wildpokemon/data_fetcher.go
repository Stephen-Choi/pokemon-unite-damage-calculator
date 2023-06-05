package wildpokemon

import (
	"encoding/json"
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"io/ioutil"
	"path/filepath"
)

const (
	wildPokemonDirectory = "/Users/stephenchoi/Desktop/Projects/pokemon-unite-damage-calculator/data/wild_pokemon_stats"
)

// fetchWildPokemonData fetches the data for a wild pokemon
func fetchWildPokemonData(wildPokemonName string, remainingTime int) (wildPokemon WildPokemon, err error) {
	WildPokemonFileName := wildPokemonName + ".json"

	// Read file
	wildPokemonFilePath := filepath.Join(wildPokemonDirectory, WildPokemonFileName)
	data, err := ioutil.ReadFile(wildPokemonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data
	var allWildPokemonStats []JsonStats
	err = json.Unmarshal(data, &allWildPokemonStats)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Convert cooldown from seconds to milliseconds
	var stats stats.Stats

	for _, wildPokemonStats := range allWildPokemonStats {
		if wildPokemonStats.RemainingTime == remainingTime {
			stats = wildPokemonStats.ToStats()
			break
		}
	}

	wildPokemon.RemainingTime = remainingTime
	wildPokemon.StartingHp = stats.Hp
	wildPokemon.Stats = stats
	return wildPokemon, nil
}
