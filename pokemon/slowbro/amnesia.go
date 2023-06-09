package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	amnesiaCooldown     = 9000.0
	amnesiaMinLevel     = 6
	amnesiaLevelUpgrade = 13
)

type Amnesia struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewAmnesia(level int) (move *Amnesia, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level < amnesiaMinLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &Amnesia{
		cooldown:   amnesiaCooldown,
		isUpgraded: level >= amnesiaLevelUpgrade,
	}
	return
}

func (move *Amnesia) GetName() string {
	return "amnesia"
}

func (move *Amnesia) CanCriticallyHit() bool {
	return false
}

func (move *Amnesia) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	// Non-damage dealing move, does not assist in ripping down an objective, so make this unavailable
	return false
}

func (move *Amnesia) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	panic("not implemented, non-damage dealing move")
}

// setLastUsed sets the lastUsed time to now
func (move *Amnesia) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
