package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	AeosCookieName         = "aeos-cookie"
	BuddyBarrierName       = "buddy-barrier"
	EnergyAmplifierName    = "energy-amplifier"
	FocusBandName          = "focus-band"
	RapidFireScarfName     = "rapid-fire-scarf"
	RockyHelmetName        = "rocky-helmet"
	ScoreShieldName        = "score-shield"
	SpecialAttackSpecsName = "special-attack-specs"
	AssaultVestName        = "assault-vest"
	ChoiceSpecsName        = "choice-specs"
	ExpShareName           = "exp-share"
	LeftoversName          = "leftovers"
	RazorClawName          = "razor-claw"
	RustedSwordName        = "rusted-sword"
	ShellBellName          = "shell-bell"
	WeaknessPolicyName     = "weakness-policy"
	AttackWeightName       = "attack-weight"
	DrainCrownName         = "drain-crown"
	FloatStoneName         = "float-stone"
	MuscleBandName         = "muscle-band"
	RescueHoodName         = "rescue-hood"
	ScopeLensName          = "scope-lens"
	SlickSpoonName         = "slick-spoon"
	WiseGlassesName        = "wise-glasses"
)

var playableHeldItems = []string{
	AeosCookieName,
	BuddyBarrierName,
	EnergyAmplifierName,
	FocusBandName,
	RapidFireScarfName,
	RockyHelmetName,
	ScoreShieldName,
	SpecialAttackSpecsName,
	AssaultVestName,
	ChoiceSpecsName,
	ExpShareName,
	LeftoversName,
	RazorClawName,
	RustedSwordName,
	ShellBellName,
	WeaknessPolicyName,
	AttackWeightName,
	DrainCrownName,
	FloatStoneName,
	MuscleBandName,
	RescueHoodName,
	ScopeLensName,
	SlickSpoonName,
	WiseGlassesName,
}

var DamageDealingItems = []string{
	ChoiceSpecsName,
	EnergyAmplifierName,
	MuscleBandName,
	RazorClawName,
	ScopeLensName,
	RapidFireScarfName,
	SlickSpoonName,
}

// IsHeldItemPlayable checks if the given held item exists in game
func IsHeldItemPlayable(heldItemName string) bool {
	for _, playableHeldItem := range playableHeldItems {
		if heldItemName == playableHeldItem {
			return true
		}
	}
	return false
}

type HeldItemEffect struct {
	Buff             stats.Buff
	Debuff           attack.Debuff
	AdditionalDamage attack.AdditionalDamage
}

type HeldItem interface {
	GetName() string
	GetStatBoosts(originalStats stats.Stats) (updatedStats stats.Stats)
	Activate(originalStats stats.Stats, elapsedTime float64, attackOption attack.Option, attackType attack.Type) (onCooldown bool, effect HeldItemEffect, err error)
}

// HeldItemSpecialEffect contains details about a special effect that a held item provides
type HeldItemSpecialEffect struct {
	Stack            *StackEffect            `json:"stack increase,omitempty"`
	AdditionalDamage *AdditionalDamageEffect `json:"additional damage,omitempty"`
	FlatStatIncrease *FlatStatIncrease       `json:"flat stat increase,omitempty"`
	Buff             *Buff                   `json:"buff,omitempty"`
}

type StackEffect struct {
	Amount int    `json:"amount"`
	Stat   string `json:"stat"`
}

type AdditionalDamageEffect struct {
	Damage           string  `json:"damage"` // Note: couldn't easily generalize, so will need to handle each damage effect separately
	Cooldown         float64 `json:"cooldown"`
	InternalCooldown float64 `json:"internal cooldown"`
	Duration         float64 `json:"duration"`
}

type FlatStatIncrease struct {
	SpecialAttack float64 `json:"sp. atk"`
}

type Buff struct {
	Effect           string  `json:"effect"`
	InternalCooldown float64 `json:"internal cooldown"`
	Duration         float64 `json:"duration"`
	Cooldown         float64 `json:"cooldown"`
}

// HeldItemData contains details about a held item
type HeldItemData struct {
	stats.Stats
	SpecialEffect HeldItemSpecialEffect `json:"special effect"`
}
