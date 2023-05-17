package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	pokemonStatsDirectory = "data/pokemon_stats"
)

// fetchPokemonStats fetches the stats of a pokemon at a given level
func fetchPokemonStats(pokemonName string, level int) (stats Stats, err error) {
	pokemonFileName := pokemonName + "_stats.json"

	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a struct
	var data MyData
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

}
