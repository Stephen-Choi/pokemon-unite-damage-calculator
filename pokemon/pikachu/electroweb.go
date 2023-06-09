package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	electrowebCooldown = 9000.0
	electrowebMaxLevel = 3
)

type Electroweb struct {
	cooldown float64
	lastUsed float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used     bool    // used is a boolean which is used to check if the move has ever been used
}

func NewElectroweb(level int) (move *Electroweb, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level > electrowebMaxLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Electroweb{
		cooldown: electrowebCooldown,
	}
	return
}

func (move *Electroweb) GetName() string {
	return "electroweb"
}

func (move *Electroweb) CanCriticallyHit() bool {
	return false
}

func (move *Electroweb) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *Electroweb) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	damage := 0.36*originalStats.SpecialAttack + 11*float64(originalStats.Level-1) + 350

	result = attack.Result{
		AttackOption:    attack.Move1,
		AttackName:      move.GetName(),
		AttackType:      attack.SpecialAttack,
		BaseDamageDealt: damage,
		NumberOfHits:    1,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Electroweb) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
