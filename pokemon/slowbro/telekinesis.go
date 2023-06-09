package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	telekinesisCooldown     = 7500.0
	telekinesisMinLevel     = 6
	telekinesisLevelUpgrade = 13
)

type Telekinesis struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewTelekinesis(level int) (move *Telekinesis, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < telekinesisMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Telekinesis{
		cooldown:   telekinesisCooldown,
		isUpgraded: level >= telekinesisLevelUpgrade,
	}
	return
}

func (move *Telekinesis) GetName() string {
	return "telekinesis"
}

func (move *Telekinesis) CanCriticallyHit() bool {
	return false
}

func (move *Telekinesis) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	// Non-damage dealing move, does not assist in ripping down an objective, so make this unavailable
	return false
}

func (move *Telekinesis) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	panic("not implemented, non-damage dealing move")
}

// setLastUsed sets the lastUsed time to now
func (move *Telekinesis) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
