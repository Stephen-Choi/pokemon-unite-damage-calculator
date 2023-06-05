package enemy

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

type Pokemon interface {
	IsWild() bool
	GetStats() stats.Stats
	GetRemainingHealth() float64
	GetMissingHealth() float64
	IsDefeated() bool
	ApplyDamage(damageTaken float64)
}
