package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	surfCooldown     = 8000.0
	surfLevelUpgrade = 11
	surfMinLevel     = 4
)

type Surf struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewSurf(level int) (move *Surf, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < surfMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Surf{
		cooldown:   surfCooldown,
		isUpgraded: level >= surfLevelUpgrade,
	}
	return
}

func (move *Surf) GetName() string {
	return "surf"
}

func (move *Surf) CanCriticallyHit() bool {
	return false
}

func (move *Surf) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *Surf) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	damage := 1.03*originalStats.SpecialAttack + 6.0*float64(originalStats.Level-1) + 210.0

	result = attack.Result{
		AttackOption:    attack.Move1,
		AttackName:      move.GetName(),
		AttackType:      attack.SpecialAttack,
		BaseDamageDealt: damage,
		NumberOfHits:    3,
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Surf) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
