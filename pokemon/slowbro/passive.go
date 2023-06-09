package slowbro

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const passiveDebuffDuration = 5000.0

type Passive struct {
	currentStacks   int
	timeOfLastStack float64
}

func NewPassive() (passive *Passive) {
	passive = &Passive{
		currentStacks: 0,
	}
	return
}

func (p *Passive) ShouldActivate(attackResult attack.Result, elapsedTime float64) bool {
	// Activate on any of water gun, surf, or scald (all move 1) or unite move
	return attackResult.AttackOption == attack.Move1 || attackResult.AttackOption == attack.UniteMove
}

func (p *Passive) Activate(pokemonStats stats.Stats, attackResult attack.Result, elapsedTime float64) (result attack.Result, err error) {
	p.clearStacks(elapsedTime)
	p.currentStacks += int(attackResult.NumberOfHits)
	p.timeOfLastStack = elapsedTime

	result.Debuffs = []attack.Debuff{
		{
			DebuffEffect: attack.DecreaseSpecialDefense,
			DebuffType:   attack.Percent,
			Stats: stats.Stats{
				SpecialDefense: 0.05 * float64(p.currentStacks),
			},
			FromPokemon: SlowbroName,
			DurationEnd: elapsedTime + passiveDebuffDuration,
		},
	}

	return
}

func (p *Passive) clearStacks(elapsedTime float64) {
	if p.timeOfLastStack+passiveDebuffDuration <= elapsedTime {
		p.currentStacks = 0
		p.timeOfLastStack = 0
	}
}
