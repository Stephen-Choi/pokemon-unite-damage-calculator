package damage_calculator

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math/rand"
	"time"
)

// ActionsLog is a map of time to action details
// Ex: Given a set of pokemon, the action logs at 100 milliseconds might look like:
// 100.0: {
//	"pikachu": "move1",
//	"cinderace": "move2",
//	"charizard": "N/A",
//}
type ActionsLog map[float64]map[string]string

type Result struct {
	TotalTime  float64 `json:"total_time"`
	ActionsLog `json:"actions_log"`
}

type DamageCalculator struct {
	attackingPokemon          map[string]pokemon.Pokemon
	attackingPokemonTeamBuff  stats.Buffs
	enemyPokemon              enemy.Pokemon
	inflictedDebuffs          []attack.Debuff
	overtimeDamage            []attack.OverTimeDamage
	timeOfNextAvailableAction map[string]float64         // timeOfNextAvailableAction is a map of pokemon name to the time at which the next action is available
	attacksThatCanCrit        map[string][]attack.Option // attacksThatCanCrit is a map of pokemon name to its attacks that can crit

	elapsedTime float64
}

func NewDamageCalculator(attackingPokemon map[string]pokemon.Pokemon, enemyPokemon enemy.Pokemon, teamBuffs stats.Buffs) *DamageCalculator {
	// Set up attacks that can crit
	attacksThatCanCrit := make(map[string][]attack.Option)
	for _, pokemon := range attackingPokemon {
		attacksThatCanCrit[pokemon.GetName()] = pokemon.GetMovesThatCanCrit()
	}

	return &DamageCalculator{
		attackingPokemon:          attackingPokemon,
		attackingPokemonTeamBuff:  teamBuffs,
		enemyPokemon:              enemyPokemon,
		attacksThatCanCrit:        attacksThatCanCrit,
		timeOfNextAvailableAction: make(map[string]float64),
	}
}

// CalculateRip calculates how long it takes to defeat an enemy pokemon
// NOTE: starting this off very simple, will always prioritize a skill move if available
// The order of priority is: unite move > move 2 > move 1 > basic attack
// TODO: set up an algorithm to determine the best course of action to achieve the most efficient rip
func (d *DamageCalculator) CalculateRip() Result {
	// CONTINUE HERE TOMORROW
	// CURRENTLY HAVE DONE:
	// enemy.ApplyDebuffs setup
	// crit setup
	// damage calculation setup (using the enemy's defensive stats)

	// Missing
	// - overtime damage
	// applying moves and setting the time of next available action
	// end return value
}

// shouldCrit checks if an attack should crit
func (d *DamageCalculator) shouldCrit(pokemonName string, attackOption attack.Option, elapsedTime float64) bool {
	// Check if the attack can crit
	canCrit := false
	for _, attackThatCanCrit := range d.attacksThatCanCrit[pokemonName] {
		if attackThatCanCrit == attackOption {
			canCrit = true
		}
	}

	// If attack cannot crit, return early
	if !canCrit {
		return false
	}

	// If attack can crit, check if it crits
	critChance := d.attackingPokemon[pokemonName].GetStats(elapsedTime).CriticalHitChance
	return rollCrit(critChance)
}

// rollCrit rolls a crit based on the crit chance
func rollCrit(critChance float64) bool {
	rand.Seed(time.Now().UnixNano())

	// Scale the probability to an integer range (0-10000)
	scaledProbability := int(critChance * 100)

	// Generate a random number between 0 and 10000
	randomNumber := rand.Intn(10001)

	// Check if the random number falls within the desired probability range
	if randomNumber <= scaledProbability {
		return true
	}
	return false
}

// elapsedTime elapses time by 1/15 of a second which is the smallest unit of time in the game
func (d *DamageCalculator) elapseTime() {
	d.elapsedTime += 66.67
}

// calculateDamageTaken calculates the damage taken by the enemy pokemon
func calculateDamageTaken(attackDamage float64, enemyStats stats.Stats, attackType attack.Type) float64 {
	var damageTaken float64
	if attackType == attack.PhysicalAttack {
		damageTaken = attackDamage * 600 / (600 + enemyStats.Defense)
	} else {
		damageTaken = attackDamage * 600 / (600 + enemyStats.SpecialDefense)
	}
	return damageTaken
}
