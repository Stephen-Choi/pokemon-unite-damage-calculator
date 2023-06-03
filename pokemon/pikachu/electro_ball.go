package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	electroBallCooldown     = 5000.0
	electroBallLevelUpgrade = 11
	electroBallMinLevel     = 4
)

type ElectroBall struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewElectroBall(level int) (move *ElectroBall, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < electroBallMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &ElectroBall{
		cooldown:   electroBallCooldown,
		isUpgraded: level >= electroBallLevelUpgrade,
	}
	return
}

func (move *ElectroBall) CanCriticallyHit() bool {
	return false
}

func (move *ElectroBall) IsAvailable(elapsedTime float64) bool {
	if !move.used {
		return true
	}
	return move.lastUsed+move.cooldown <= elapsedTime
}

func (move *ElectroBall) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	var damage float64
	var executionPercent float64
	if !move.isUpgraded {
		damage = 0.66*originalStats.SpecialAttack + 25*float64(originalStats.Level-1) + 530
		executionPercent = 0.04
	} else {
		damage = 0.77*originalStats.SpecialAttack + 29*float64(originalStats.Level-1) + 640
		executionPercent = 0.05
	}

	result = attack.Result{
		AttackOption: attack.Move1,
		AttackType:   attack.SpecialAttack,
		DamageDealt:  damage,
		ExecutionPercentDamage: attack.ExecutePercentDamage{
			Percent:      executionPercent,
			CappedDamage: 1200,
		},
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *ElectroBall) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
