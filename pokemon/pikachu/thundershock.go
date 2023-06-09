package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	pokemonErrors "github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	thunderShockCooldown = 5000
	thunderShockMaxLevel = 3
)

type ThunderShock struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewThunderShock(level int) (move *ThunderShock, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level > thunderShockMaxLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &ThunderShock{
		cooldown: thunderShockCooldown,
	}
	return
}

func (move *ThunderShock) GetName() string {
	return "thunder shock"
}

func (move *ThunderShock) CanCriticallyHit() bool {
	return false
}

func (move *ThunderShock) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *ThunderShock) Activate(stats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	damage := 0.75*stats.SpecialAttack + 21*float64(stats.Level-1) + 390

	result = attack.Result{
		AttackOption:    attack.Move2,
		AttackName:      move.GetName(),
		AttackType:      attack.SpecialAttack,
		BaseDamageDealt: damage,
		NumberOfHits:    1,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *ThunderShock) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
