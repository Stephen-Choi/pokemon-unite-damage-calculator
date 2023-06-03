package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	pokemonErrors "github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	voltTackleCooldown     = 10000.0
	voltTackleLevelUpgrade = 13
	voltTackleMinLevel     = 6
)

type VoltTackle struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewVoltTackle(level int) (move *VoltTackle, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < voltTackleMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	toSetVoltTackleCooldown := voltTackleCooldown
	if level >= voltTackleLevelUpgrade {
		toSetVoltTackleCooldown = voltTackleCooldown - 2000.0
	}

	move = &VoltTackle{
		cooldown:   toSetVoltTackleCooldown,
		isUpgraded: level >= voltTackleLevelUpgrade,
	}
	return
}

func (move *VoltTackle) CanCriticallyHit() bool {
	return false
}

func (move *VoltTackle) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *VoltTackle) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	damage := 0.14*originalStats.SpecialAttack + 3*float64(originalStats.Level-1) + 140

	result = attack.Result{
		AttackOption: attack.Move2,
		AttackType:   attack.SpecialAttack,
		DamageDealt:  damage,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *VoltTackle) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
