package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math"
)

const (
	thunderCooldown     = 8000.0
	thunderLevelUpgrade = 11
	thunderMinLevel     = 4
)

type Thunder struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewThunder(level int) (move *Thunder, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < thunderMinLevel {
		err = errors.New(pokemon.ErrInvalidMovesetForLevel)
		return
	}

	move = &Thunder{
		cooldown:   thunderCooldown,
		isUpgraded: level >= thunderLevelUpgrade,
	}
	return
}

func (move *Thunder) IsAvailable(elapsedTime float64) bool {
	if !move.used {
		return true
	}
	return move.lastUsed+move.cooldown <= elapsedTime
}

func (move *Thunder) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	// Damage calculation
	damagePerHit := 0.2*originalStats.SpecialAttack + 5*float64(originalStats.Level-1) + 210
	damagePerHit = math.Floor(damagePerHit*100) / 100

	// Overtime damage calculation
	var damageFrequency float64
	moveDuration := 2500.0
	if !move.isUpgraded {
		numThunderStrikes := 5.0
		damageFrequency = moveDuration / numThunderStrikes
	} else {
		numThunderStrikes := 7.0
		damageFrequency = moveDuration / numThunderStrikes
	}

	result = attack.Result{
		AttackOption: attack.Move1,
		AttackType:   attack.SpecialAttack,
		OvertimeDamage: attack.OverTimeDamage{
			Damage:          damagePerHit,
			DamageFrequency: damageFrequency,
			DurationEnd:     elapsedTime + moveDuration,
		},
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Thunder) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
