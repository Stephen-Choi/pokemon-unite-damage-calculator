package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	thunderboltCooldown     = 8000.0
	thunderboltLevelUpgrade = 13
	thunderboltMinLevel     = 6
)

type Thunderbolt struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewThunderbolt(level int) (move *Thunderbolt, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < thunderboltMinLevel {
		err = errors.New(pokemon.ErrInvalidMovesetForLevel)
		return
	}

	move = &Thunderbolt{
		cooldown:   thunderboltCooldown,
		isUpgraded: level >= thunderboltLevelUpgrade,
	}
	return
}

func (move *Thunderbolt) IsAvailable(elapsedTime float64) bool {
	if !move.used {
		return true
	}
	return move.lastUsed+move.cooldown <= elapsedTime
}

func (move *Thunderbolt) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	var damage float64
	if !move.isUpgraded {
		damage = 0.50*originalStats.SpecialAttack + 12*float64(originalStats.Level-1) + 500
	} else {
		damage = 0.59*originalStats.SpecialAttack + 14*float64(originalStats.Level-1) + 600
	}

	result = attack.Result{
		AttackOption: attack.Move2,
		AttackType:   attack.SpecialAttack,
		DamageDealt:  damage,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Thunderbolt) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
