package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	scaldCooldown     = 5000.0
	scaldLevelUpgrade = 11
	scaldMinLevel     = 4
)

type Scald struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewScald(level int) (move *Scald, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < scaldMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Scald{
		cooldown:   scaldCooldown,
		isUpgraded: level >= scaldLevelUpgrade,
	}
	return
}

func (move *Scald) GetName() string {
	return "scald"
}

func (move *Scald) CanCriticallyHit() bool {
	return false
}

func (move *Scald) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *Scald) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	maxNumHits := 3.0
	maxNumburns := 5.0
	damagePerHit := 1.0*originalStats.SpecialAttack + 7*float64(originalStats.Level-1) + 160
	damagePerBurn := 0.2*originalStats.SpecialAttack + 1*float64(originalStats.Level-1) + 32
	totalDamage := damagePerHit*maxNumHits + damagePerBurn*maxNumburns

	result = attack.Result{
		AttackOption:    attack.Move1,
		AttackName:      move.GetName(),
		AttackType:      attack.SpecialAttack,
		BaseDamageDealt: totalDamage,
		NumberOfHits:    maxNumHits,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Scald) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
