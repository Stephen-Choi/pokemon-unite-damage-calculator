package battleitems

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"

type NonDamagingItem struct{}

func (item NonDamagingItem) Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, buff stats.Buff, err error) {
	return // do nothing
}

func (item NonDamagingItem) IsAvailable(elapsedTime float64) bool {
	return false // non damage related so not relevant for damage calculations, simply return false
}
