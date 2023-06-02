package enemy

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

type Pokemon interface {
	IsWild() bool
	GetStats() stats.Stats
}

type DefaultEnemy struct {
	Wild  bool
	stats stats.Stats
}

func (enemy DefaultEnemy) IsWild() bool {
	return enemy.Wild
}

func (enemy DefaultEnemy) GetStats() stats.Stats {
	return enemy.stats
}

// ApplyDebuffs applies debuffs to the enemy pokemon for a specific attacking pokemon
// This needs to occur on a per pokemon basis since a pokemon can have debuff effects that only apply for itself (ex. slick spoon)
func ApplyDebuffs(attackingPokemon string, enemy Pokemon, debuffs []attack.Debuff) (enemyStats stats.Stats) {
	enemyStats = enemy.GetStats()

	// Debuffs apply additively
	debuffStats := stats.Stats{}

	for _, debuff := range debuffs {
		switch debuff.DebuffEffect {
		case attack.IgnoreDefenseForAttackingPokemon:
			if attackingPokemon == debuff.FromPokemon {
				debuffStats.SpecialDefense += debuff.Stats.SpecialDefense
			}
		}
	}

	enemyStats.Hp *= 1 - debuffStats.Hp
	enemyStats.Attack *= 1 - debuffStats.Attack
	enemyStats.Defense *= 1 - debuffStats.Defense
	enemyStats.SpecialAttack *= 1 - debuffStats.SpecialAttack
	enemyStats.SpecialDefense *= 1 - debuffStats.SpecialDefense
	enemyStats.AttackSpeed *= 1 - debuffStats.AttackSpeed
	enemyStats.CriticalHitChance *= 1 - debuffStats.CriticalHitChance
	enemyStats.CriticalHitDamage *= 1 - debuffStats.CriticalHitDamage
	enemyStats.CooldownReduction *= 1 - debuffStats.CooldownReduction
	enemyStats.EnergyRate *= 1 - debuffStats.EnergyRate
	return enemyStats
}
