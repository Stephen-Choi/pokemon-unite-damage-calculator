package pikachu

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	thunderstormCooldown = 80000.0
	thunderstormMinLevel = 9
)

type Thunderstorm struct {
	pokemonLevel int
	cooldown     float64
	lastUsed     float64 // lastUsed is a time in milliseconds which is used to check if the move is on cooldown
	used         bool    // used is a boolean which is used to check if the move has ever been used
}

func NewThunderstorm(level int) (move *Thunderstorm) {
	move = &Thunderstorm{
		pokemonLevel: level,
		cooldown:     thunderstormCooldown,
	}
	return
}

func (move *Thunderstorm) IsAvailable(elapsedTime float64) bool {
	// Unite move is unlocked at level 9
	if move.pokemonLevel < thunderstormMinLevel {
		return false
	}

	if !move.used {
		return true
	}
	return move.lastUsed+move.cooldown <= elapsedTime
}

func (move *Thunderstorm) Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) { // Damage calculation
	var damagePerHit float64
	isEnemyWildPokemon := enemyPokemon.IsWild()
	if isEnemyWildPokemon {
		damagePerHit = 0.49*originalStats.SpecialAttack + 10*float64(originalStats.Level-1) + 490
	} else {
		damagePerHit = 0 // Unite move only hits enemy players
	}

	// Overtime damage calculation
	moveDuration := 3000.0
	numThunderStrikes := 4
	damageFrequency := moveDuration / float64(numThunderStrikes)

	uniteBuffDuration := 6000.0

	result = attack.Result{
		OvertimeDamage: attack.OverTimeDamage{
			Damage:          damagePerHit,
			DamageFrequency: damageFrequency,
			DurationEnd:     elapsedTime + moveDuration,
		},
		Buff: stats.StatBuff{
			DurationEnd: uniteBuffDuration + elapsedTime,
			Stats: stats.Stats{
				CooldownReduction: 0.3,
			},
			BuffType: stats.PercentIncrease,
		},
	}
	move.setLastUsed(elapsedTime)
	return
}

// setLastUsed sets the lastUsed time to now
func (move *Thunderstorm) setLastUsed(elapsedTime float64) {
	move.lastUsed = elapsedTime
	move.used = true
}