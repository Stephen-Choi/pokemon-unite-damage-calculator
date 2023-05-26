package pikachu

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const stacksNeededForBoostedBasicAttack = 2

type BasicAttack struct {
	boostedStack       int     // boostedStack is the number of stacks leading to a boosted basic attack
	boostedStackExpiry float64 // boostedStackExpiry is the time in milliseconds when the boostedStack will expire
}

func NewBasicAttack() (basicAttack *BasicAttack) {
	basicAttack = &BasicAttack{}
	return
}

func (ba *BasicAttack) Attack(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	if ba.boostedStackExpiry <= elapsedTime {
		ba.resetStack()
	}

	var damage float64
	if ba.boostedStack >= stacksNeededForBoostedBasicAttack {
		damage = 0.38*originalStats.SpecialAttack + 10*float64(originalStats.Level-1) + 200
		ba.resetStack() // reset boosted stack after using it
	} else {
		damage = 1.00 * originalStats.Attack
		ba.boostedStack++                                                             // add boosted stack
		ba.boostedStackExpiry = elapsedTime + attack.BoostedStackDurationBeforeExpiry // update the boosted stack expiry
	}

	result = attack.Result{
		DamageDealt: damage,
	}
	return
}

func (ba *BasicAttack) resetStack() {
	ba.boostedStack = 0
}
