package damage_calculator

import (
	"fmt"
	enemy2 "github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
	"github.com/stretchr/testify/assert"
	"testing"
)

// setupDamageCalculatorScenario sets up a damage calculator scenario for testing
func setupDamageCalculatorScenario() (attackingPokemon map[string]pokemon.Pokemon, enemy enemy2.Pokemon) {
	// Set up attacking pokemon
	battleItem, err := battleitems.GetBattleItem(battleitems.EjectButton)
	if err != nil {
		panic(fmt.Sprintf("error getting battle item %s: %s", battleitems.EjectButton, err))
	}

	pikachu, err := pokemon.GetPokemon(pokemon.PikachuName, 15, "electro_ball", "volt_tackle", nil, battleItem)
	if err != nil {
		panic(err)
	}

	attackingPokemon = make(map[string]pokemon.Pokemon)
	attackingPokemon[pokemon.PikachuName] = pikachu

	// Set up enemy
	enemy = &enemy2.DefaultEnemy{
		Wild: true,
		Stats: stats.Stats{
			Hp:             12350, // Regis stats
			Defense:        250,
			SpecialDefense: 250,
		},
	}

	return attackingPokemon, enemy
}

func Test_NewDamageCalculator(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		attackingPokemon, enemy := setupDamageCalculatorScenario()
		damageCalc := NewDamageCalculator(attackingPokemon, enemy, nil)

		// Assert damage calculator is setup as expected
		assert.Equal(t, 1, len(damageCalc.attackingPokemon))
		assert.Equal(t, attackingPokemon[pokemon.PikachuName], damageCalc.attackingPokemon[pokemon.PikachuName])
		assert.Equal(t, enemy, damageCalc.enemyPokemon)
		assert.Equal(t, 0, len(damageCalc.inflictedDebuffs))
		assert.Equal(t, 0, len(damageCalc.overtimeDamageByPokemon))
		assert.Equal(t, 0, len(damageCalc.timeOfNextAvailableAction))
	})
}

func Test_CalculateRip(t *testing.T) {
	t.Run("simple rip with one attacking pokemon", func(t *testing.T) {
		attackingPokemon, enemy := setupDamageCalculatorScenario()
		damageCalc := NewDamageCalculator(attackingPokemon, enemy, nil)

		// Test rip calculation works
		result, err := damageCalc.CalculateRip()

		fmt.Println("total time: ", result.TotalTime)
		fmt.Println("num of state log: ", len(result.StateLog))

		assert.NoError(t, err)

		// Assert result is within expectations
		// For a lvl 15 pikachu, should be able to defeat a 7:00 regi in the 20 seconds range
		assert.True(t, result.TotalTime > time.ConvertSecondsToMilliseconds(10) && result.TotalTime < time.ConvertSecondsToMilliseconds(30))
		// Since it should take at least 10 seconds, there should be at least 150 actions (10 seconds * 15 action states per second)
		assert.True(t, len(result.StateLog) > 150)
	})
}
