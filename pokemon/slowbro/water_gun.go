package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemonErrors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	waterGunCooldown     = 5000.0
	waterGunLevelUpgrade = 11
	waterGunMaxLevel     = 4
)

type WaterGun struct {
	cooldown   float64
	isUpgraded bool    // isUpgraded is a boolean which is used to check if the move has been upgraded
	lastUsed   float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used       bool    // used is a boolean which is used to check if the move has ever been used
}

func NewWaterGun(level int) (move *WaterGun, err error) {
	// Ensure moveset is valid for the current pokemon level
	if level >= waterGunMaxLevel {
		err = errors.New(pokemonErrors.ErrInvalidMovesetForLevel)
		return
	}

	move = &WaterGun{
		cooldown:   waterGunCooldown,
		isUpgraded: level >= waterGunLevelUpgrade,
	}
	return
}

func (move *WaterGun) GetName() string {
	return "water gun"
}

func (move *WaterGun) CanCriticallyHit() bool {
	return false
}

func (move *WaterGun) IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool {
	if !move.used {
		return true
	}
	// Apply cooldown reduction
	updatedCooldown := move.cooldown * (1 - pokemonStats.CooldownReduction)

	return move.lastUsed+updatedCooldown <= elapsedTime
}

func (move *WaterGun) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	damage := 2.94*originalStats.SpecialAttack + 26.0*float64(originalStats.Level-1) + 480.0

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
func (move *WaterGun) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}
