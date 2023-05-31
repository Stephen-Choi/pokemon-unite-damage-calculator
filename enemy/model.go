package enemy

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"

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
