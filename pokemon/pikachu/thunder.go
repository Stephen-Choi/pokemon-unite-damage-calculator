package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
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
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Thunder{
		cooldown:   thunderCooldown,
		isUpgraded: level >= thunderLevelUpgrade,
	}
	return
}

func (move *Thunder) GetName() string {
	return "thunder"
}

func (move *Thunder) CanCriticallyHit() bool {
	return false
}

func (move *Thunder) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *Thunder) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	// BaseDamage calculation
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
		AttackOption:    attack.Move1,
		AttackName:      move.GetName(),
		AttackType:      attack.SpecialAttack,
		BaseDamageDealt: damagePerHit, // Deal damage for this first hit
		OvertimeDamage: attack.OverTimeDamage{
			Source:                  move.GetName(),
			AttackType:              attack.SpecialAttack,
			BaseDamage:              damagePerHit,
			LastInflictedDamageTime: elapsedTime,
			DamageFrequency:         damageFrequency,
			DurationStart:           elapsedTime,
			DurationEnd:             elapsedTime + moveDuration,
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
