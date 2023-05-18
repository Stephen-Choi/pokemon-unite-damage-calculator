package helditems

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	AeosCookieName         = "aeosCookie"
	BuddyBarrierName       = "buddyBarrier"
	EnergyAmplifierName    = "energyAmplifier"
	FocusBandName          = "focusBand"
	RapidFireScarfName     = "rapidFireScarf"
	RockyHelmetName        = "rockyHelmet"
	ScoreShieldName        = "scoreShield"
	SpecialAttackSpecsName = "specialAttackSpecs"
	AssaultVestName        = "assaultVest"
	ChoiceSpecsName        = "choiceSpecs"
	ExpShareName           = "expShare"
	LeftoversName          = "leftovers"
	RazorClawName          = "razorClaw"
	RustedSwordName        = "rustedSword"
	ShellBellName          = "shellBell"
	WeaknessPolicyName     = "weaknessPolicy"
	AttackWeightName       = "attackWeight"
	DrainCrownName         = "drainCrown"
	FloatStoneName         = "floatStone"
	MuscleBandName         = "muscleBand"
	RescueHoodName         = "rescueHood"
	ScopeLensName          = "scopeLens"
	SlickSpoonName         = "slickSpoon"
	WiseGlassesName        = "wiseGlasses"
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

// IsHeldItemPlayable checks if the given held item exists in game
func IsHeldItemPlayable(heldItemName string) bool {
	for _, playableHeldItem := range playableHeldItems {
		if heldItemName == playableHeldItem {
			return true
		}
	}
	return false
}

type HeldItemEffects struct {
	StatBoosts      stats.Stats
	StatusCondition []attack.StatusConditions
}

type HeldItem interface {
	GetStatBoosts() stats.Stats
	Activate() (onCooldown bool, effect HeldItemEffects, err error)
}

// HeldItemSpecialEffect contains details about a special effect that a held item provides
type HeldItemSpecialEffect struct {
	Stack            StackEffect            `json:"stack increase"`
	AdditionalDamage AdditionalDamageEffect `json:"additional damage"`
	FlatStatIncrease FlatStatIncrease       `json:"flat stat increase"`
	Buff             Buff                   `json:"buff"`
}

type StackEffect struct {
	Amount int    `json:"amount"`
	Stat   string `json:"stat"`
}

type AdditionalDamageEffect struct {
	AdditionalDamage string  `json:"additional damage"` // Note: couldn't easily generalize, so will need to handle each damage effect separately
	AdditionalNote   string  `json:"additional note"`
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
	Stats         stats.JsonStats
	SpecialEffect HeldItemSpecialEffect `json:"special effect"`
}
