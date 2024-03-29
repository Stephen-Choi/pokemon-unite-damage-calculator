package enemy

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"

type DefaultEnemy struct {
	Wild       bool
	StartingHP float64
	Stats      stats.Stats
}

func (enemy *DefaultEnemy) IsWild() bool {
	return enemy.Wild
}

func (enemy *DefaultEnemy) GetStats() stats.Stats {
	return enemy.Stats
}

func (enemy *DefaultEnemy) IsDefeated() bool {
	return enemy.Stats.Hp <= 0
}

func (enemy *DefaultEnemy) GetRemainingHealth() float64 {
	return enemy.Stats.Hp
}

func (enemy *DefaultEnemy) GetMissingHealth() float64 {
	return enemy.StartingHP - enemy.Stats.Hp
}

func (enemy *DefaultEnemy) ApplyDamage(damageTaken float64) {
	enemy.Stats.Hp = enemy.Stats.Hp - damageTaken
}
