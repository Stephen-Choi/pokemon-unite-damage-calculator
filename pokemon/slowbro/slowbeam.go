package slowbro

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	slowbeamCooldown = 100000.0
	slowbeamMinLevel = 9
)

type Slowbeam struct {
	cooldown     float64
	pokemonLevel int
	lastUsed     float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used         bool    // used is a boolean which is used to check if the move has ever been used
}

func NewSlowbeam(level int) (move *Slowbeam) {
	move = &Slowbeam{
		pokemonLevel: level,
		cooldown:     slowbeamCooldown,
	}
	return
}

func (move *Slowbeam) GetName() string {
	return "slowbeam (unite move)"
}

func (move *Slowbeam) CanCriticallyHit() bool {
	return false
}

func (move *Slowbeam) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	// Can only hit enemy players, not needed for wild pokemon damage calculation
	return false
}

func (move *Slowbeam) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	panic("not implemented, non-damage dealing move")
}

// setLastUsed sets the lastUsed time to now
func (move *Slowbeam) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
