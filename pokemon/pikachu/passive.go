package pikachu

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

type Passive struct{}

func NewPassive() (passive *Passive) {
	passive = &Passive{}
	return
}

func (p *Passive) IsAvailable(elapsedTime float64) bool {
	return false
}

func (p *Passive) Activate(pokemonStats stats.Stats, attackResult attack.Result, elapsedTime float64) (result attack.Result, err error) {
	panic("passive is not implemented (not needed for damage calc)")
}
