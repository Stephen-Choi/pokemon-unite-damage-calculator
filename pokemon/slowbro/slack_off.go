package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	slackOffCooldown = 11000.0
	slackOffMaxLevel = 6
)

type SlackOff struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewSlackOff(level int) (move *SlackOff, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level >= slackOffMaxLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &SlackOff{
		cooldown: slackOffCooldown,
	}
	return
}

func (move *SlackOff) GetName() string {
	return "slack off"
}

func (move *SlackOff) CanCriticallyHit() bool {
	return false
}

func (move *SlackOff) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	// Non damage dealing move, does not assist in ripping down an objectvie so make this unavailable
	return false
}

func (move *SlackOff) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	panic("not implemented, non damage dealing move")
}

// setLastUsed sets the lastUsed time to now
func (move *SlackOff) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
