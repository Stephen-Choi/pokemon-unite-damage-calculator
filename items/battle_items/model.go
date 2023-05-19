package battleitems

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"

const (
	FluffyTailName = "fluffy_tail"
	XAttackName    = "x_attack"
)

var playableBattleItems = []string{
	FluffyTailName,
	XAttackName,
}

// IsBattleItemPlayable checks if the given battle item exists in game
func IsBattleItemPlayable(battleItemName string) bool {
	for _, playableBattleItem := range playableBattleItems {
		if battleItemName == playableBattleItem {
			return true
		}
	}
	return false
}

type BattleItemEffect struct {
	UpdatedStats     *stats.Stats
	AdditionalDamage *float64
}

type BattleItem interface {
	Activate(originalStats stats.Stats) (onCooldown bool, effect BattleItemEffect, err error)
}

type StatsBuff struct {
	AttackBuff        float64 `json:"attack,omitempty"`
	AttackSpeedBuff   float64 `json:"atk speed,omitempty"`
	SpecialAttackBuff float64 `json:"sp. atk,omitempty"`
}

type BattleItemSpecialEffect struct {
	AdditionalDamageEffect *string    `json:"additional damage,omitempty"`
	StatsBuff              *StatsBuff `json:"stats buff,omitempty"`
}

type BattleItemData struct {
	Cooldown      int                     `json:"cooldown"`
	SpecialEffect BattleItemSpecialEffect `json:"special effect"`
}
