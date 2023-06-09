package battleitems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	FluffyTailName = "fluffy-tail"
	XAttackName    = "x-attack"
	EjectButton    = "eject-button"
	FullHeal       = "full-heal"
	GoalGetter     = "goal-getter"
	Potion         = "potion"
	SlowSmoke      = "slow-smoke"
	XSpeed         = "x-speed"
)

var playableBattleItems = []string{
	FluffyTailName,
	XAttackName,
	EjectButton,
	FullHeal,
	GoalGetter,
	Potion,
	SlowSmoke,
	XSpeed,
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
	Buff             stats.Buff
	AdditionalDamage attack.AdditionalDamage
}

type BattleItem interface {
	GetName() string
	IsAvailable(elapsedTime float64) bool
	Activate(originalStats stats.Stats, elapsedTime float64) (onCooldown bool, battleItemEffect BattleItemEffect, err error)
}

type StatsBuff struct {
	AttackBuff        float64 `json:"attack,omitempty"`
	AttackSpeedBuff   float64 `json:"atk spd,omitempty"`
	SpecialAttackBuff float64 `json:"sp. attack,omitempty"`
	Duration          float64 `json:"duration,omitempty"`
}

type AdditionalDamageEffect struct {
	Damage   float64
	Duration float64
}

type AdditionalDamageEffectJSON struct {
	Damage   string  `json:"-"` // Note: couldn't easily generalize, so will need to handle each damage effect separately
	Duration float64 `json:"duration"`
}

type BattleItemSpecialEffect struct {
	AdditionalDamage AdditionalDamageEffectJSON `json:"additional damage,omitempty"`
	StatsBuff        StatsBuff                  `json:"stats buff,omitempty"`
}

type BattleItemData struct {
	Cooldown      float64                 `json:"cooldown"` // Cooldown time in milliseconds
	SpecialEffect BattleItemSpecialEffect `json:"special effect"`
}
