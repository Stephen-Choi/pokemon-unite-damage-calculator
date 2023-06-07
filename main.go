package main

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/damagecalculator"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy/wildpokemon"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/calculate-rip", HandleCalculateRip)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type PokemonSetup struct {
	Name       string          `json:"name"`
	Level      int             `json:"level"`
	Move1      string          `json:"move1"`
	Move2      string          `json:"move2"`
	HeldItems  []HeldItemSetup `json:"held-items"`
	BattleItem string          `json:"battle-item"`
}

type HeldItemSetup struct {
	Name   string `json:"name"`
	Stacks *int   `json:"stacks,omitempty"`
}

type EnemySetup struct {
	Name  string `json:"name"`
	Level *int   `json:"level,omitempty"` // Required if setting an enemy pokemon (NOTE: currently not available)
}

type DamageCalculationRequest struct {
	Team          []PokemonSetup `json:"team"`
	Enemy         EnemySetup     `json:"enemy"`
	TeamBuffs     []string       `json:"team-buffs"`
	TimeRemaining int            `json:"time-remaining"`
}

func HandleCalculateRip(c *gin.Context) {
	var damageCalculationRequest DamageCalculationRequest

	// Bind the request body to the DamageCalculationRequest struct
	if err := c.ShouldBindJSON(&damageCalculationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attackingPokemon := make(map[string]pokemon.Pokemon)
	// Setup attacking pokemon team
	for _, pokemonSetup := range damageCalculationRequest.Team {
		// Setup held items
		heldItems := []helditems.HeldItem{}
		for _, heldItemSetup := range pokemonSetup.HeldItems {
			heldItem, err := helditems.GetHeldItem(heldItemSetup.Name, heldItemSetup.Stacks)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			heldItems = append(heldItems, heldItem)
		}

		// Setup battle item
		battleItem, err := battleitems.GetBattleItem(pokemonSetup.BattleItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Setup pokemon
		pokemon, err := pokemon.GetPokemon(pokemonSetup.Name, pokemonSetup.Level, pokemonSetup.Move1, pokemonSetup.Move2, heldItems, battleItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		attackingPokemon[pokemonSetup.Name] = pokemon
	}

	// Setup enemy
	// NOTE: currently only wild pokemon are available for selection
	enemy, err := wildpokemon.GetObjectivePokemon(damageCalculationRequest.Enemy.Name, damageCalculationRequest.TimeRemaining)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Setup team buffs
	teamBuffs, err := stats.GetTeamBuffs(damageCalculationRequest.TeamBuffs, damageCalculationRequest.TimeRemaining)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Setup damage calculator
	damageCalculator := damagecalculator.NewDamageCalculator(attackingPokemon, enemy, teamBuffs)

	// Calculate damage
	response, err := damageCalculator.CalculateRip()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
