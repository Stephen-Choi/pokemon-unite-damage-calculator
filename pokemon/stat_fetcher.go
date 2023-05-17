package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

const (
	pokemonStatsDirectory = "data/pokemon_stats"
)

// JsonStats is a temporary struct used to unmarshal the pokemon json stats
type JsonStats struct {
	Level             string `json:"level"`
	Hp                string `json:"hp"`
	Attack            string `json:"attack"`
	Defense           string `json:"def"`
	SpecialAttack     string `json:"sp. attack"`
	SpecialDefense    string `json:"sp. def"`
	AttackSpeed       string `json:"atk spd"`
	CriticalHitChance string `json:"crit chance"`
	CooldownReduction string `json:"CDR"`
}

// fetchPokemonStats fetches the stats of a pokemon at a given level
func fetchPokemonStats(pokemonName string, level int) (stats Stats, err error) {
	pokemonFileName := pokemonName + "_stats.json"

	// Read file
	pokemonFilePath := filepath.Join(pokemonStatsDirectory, pokemonFileName)
	data, err := ioutil.ReadFile(pokemonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Pokemon JsonStats
	var pokemonJsonStats []JsonStats
	err = json.Unmarshal(data, &pokemonJsonStats)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Convert the json stats for the specified level to internal use struct (with better typing)
	index := level - 1 // level 1 is at index 0
	return toTypedStats(pokemonJsonStats[index])
}

// toTypedStats converts the JsonStats struct to the internal use typed Stats struct
func toTypedStats(jsonStats JsonStats) (Stats, error) {
	level, err := strconv.Atoi(jsonStats.Level)
	if err != nil {
		fmt.Println("Error parsing level:", err)
		return Stats{}, err
	}
	hp, err := strconv.Atoi(jsonStats.Hp)
	if err != nil {
		fmt.Println("Error parsing hp:", err)
		return Stats{}, err
	}
	attack, err := strconv.Atoi(jsonStats.Attack)
	if err != nil {
		fmt.Println("Error parsing attack:", err)
		return Stats{}, err
	}
	defense, err := strconv.Atoi(jsonStats.Defense)
	if err != nil {
		fmt.Println("Error parsing defense:", err)
		return Stats{}, err
	}
	specialAttack, err := strconv.Atoi(jsonStats.SpecialAttack)
	if err != nil {
		fmt.Println("Error parsing special attack:", err)
		return Stats{}, err
	}
	specialDefense, err := strconv.Atoi(jsonStats.SpecialDefense)
	if err != nil {
		fmt.Println("Error parsing special defense:", err)
		return Stats{}, err
	}
	attackSpeed, err := convertPercentToFloat(jsonStats.AttackSpeed)
	if err != nil {
		fmt.Println("Error parsing attack speed:", err)
		return Stats{}, err
	}
	criticalHitChance, err := convertPercentToFloat(jsonStats.CriticalHitChance)
	if err != nil {
		fmt.Println("Error parsing critical hit chance:", err)
		return Stats{}, err
	}
	cooldownReduction, err := convertPercentToFloat(jsonStats.CooldownReduction)
	if err != nil {
		fmt.Println("Error parsing cooldown reduction:", err)
		return Stats{}, err
	}

	return Stats{
		level:             level,
		hp:                hp,
		attack:            attack,
		defense:           defense,
		specialAttack:     specialAttack,
		specialDefense:    specialDefense,
		attackSpeed:       attackSpeed,
		criticalHitChance: criticalHitChance,
		criticalHitDamage: 2.0,
		cooldownReduction: cooldownReduction,
	}, nil
}

func convertPercentToFloat(percentString string) (float64, error) {
	percentString = percentString[:len(percentString)-1] // Remove the '%' symbol
	percentValue, err := strconv.ParseFloat(percentString, 64)
	if err != nil {
		fmt.Println("Error parsing percent:", err)
		return 0, err
	}

	// Convert percent to a decimal value
	decimalValue := percentValue / 100.0
	return decimalValue, nil
}
