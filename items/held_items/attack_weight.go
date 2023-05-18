package helditems

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"

// AttackWeight is a held item that increases attack stats based on stacks
type AttackWeight struct {
	Stacks int
}

func NewAttackWeight(stacks int) *AttackWeight {
	return &AttackWeight{
		Stacks: stacks,
	}
}

func (a *AttackWeight) GetStatBoosts() stats.Stats {
	return stats.Stats{}
}
