package attack

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/samber/lo"
	"math"
)

type Type string

const (
	PhysicalAttack Type = "physical"
	SpecialAttack  Type = "special"
)

// Option is an enum for the different types of attacks a pokemon can do
type Option string

const (
	Move1                  Option = "move1"
	Move2                  Option = "move2"
	UniteMove              Option = "uniteMove"
	BasicAttackOption      Option = "basicAttack"
	CriticalHitBasicAttack Option = "criticalHitBasicAttack"
)

// DebuffEffect is an enum for the different types of status conditions a pokemon can inflict
type DebuffEffect string

const (
	IgnoreDefenseForAttackingPokemon DebuffEffect = "ignoreDefense"
)

type DebuffType string

const (
	Percent DebuffType = "percent"
	Flat    DebuffType = "flat"
)

// Debuff is a struct containing the type and amount of debuff to be applied
type Debuff struct {
	DebuffEffect DebuffEffect
	DebuffType   DebuffType
	stats.Stats
	FromPokemon string // fromPokemon is the name of the pokemon that applied the debuff
	DurationEnd float64
}

func (d Debuff) Exists() bool {
	return d != Debuff{}
}

type AdditionalDamageType string

const (
	SimpleAdditionalDamage AdditionalDamageType = "single"             // SingleAdditionalDamage is a single instance of additional damage
	PercentDamageBoost     AdditionalDamageType = "percentDamageBoost" // PercentDamageBoost is a damage boost (as a percent increase) that lasts for a certain amount of time
	RemainingEnemyHp       AdditionalDamageType = "remainingEnemyHp"   // RemainingEnemyHp is additional damage that scales with the enemy's remaining HP
)

// AdditionalDamageName is the name of the origin of the additional damage to be applied for damage dealing moves
type AdditionalDamageName string

const (
	Move1AdditionalDamage       AdditionalDamageName = "move1AdditionalDamage"
	Move2AdditionalDamage       AdditionalDamageName = "move2AdditionalDamage"
	UniteMoveAdditionalDamage   AdditionalDamageName = "uniteMoveAdditionalDamage"
	BasicAttackAdditionalDamage AdditionalDamageName = "basicAttackAdditionalDamage"
	BattleItemAdditionalDamage  AdditionalDamageName = "battleItemAdditionalDamage"
	HeldItem1AdditionalDamage   AdditionalDamageName = "heldItem1AdditionalDamage"
	HeldItem2AdditionalDamage   AdditionalDamageName = "heldItem2AdditionalDamage"
	HeldItem3AdditionalDamage   AdditionalDamageName = "heldItem3AdditionalDamage"
)

func GetAdditionalDamageName(index int) AdditionalDamageName {
	switch index {
	case 0:
		return HeldItem1AdditionalDamage
	case 1:
		return HeldItem2AdditionalDamage
	case 2:
		return HeldItem3AdditionalDamage
	default:
		return ""
	}
}

// AdditionalDamage is a struct containing the type and amount of additional damage to be applied
type AdditionalDamage struct {
	Type         AdditionalDamageType
	Amount       float64
	CappedAmount *float64 // only applicable to certain held items (i.e muscle band)
	DurationEnd  *float64 // only applicable to certain held items (i.e energy amplifier)
}

func (a AdditionalDamage) Exists() bool {
	return a != AdditionalDamage{}
}

type AllAdditionalDamage map[AdditionalDamageName]AdditionalDamage

func NewAllAdditionalDamage() AllAdditionalDamage {
	return make(AllAdditionalDamage)
}

func (a AllAdditionalDamage) Add(additionalDamageName AdditionalDamageName, additionalDamage AdditionalDamage) {
	a[additionalDamageName] = additionalDamage
}

// Calculate calculates the total additional damage
func (a AllAdditionalDamage) Calculate(baseAttackDamage float64, enemyStats stats.Stats) float64 {
	// Damage must occur for additional damage boosts to be applied
	if baseAttackDamage == 0.0 {
		return 0.0
	}

	totalAdditionalDamage := 0.0
	for _, additionalDamage := range a {
		switch additionalDamage.Type {
		case SimpleAdditionalDamage:
			totalAdditionalDamage += additionalDamage.Amount
		case PercentDamageBoost:
			totalAdditionalDamage += baseAttackDamage * additionalDamage.Amount
		case RemainingEnemyHp:
			if additionalDamage.CappedAmount == nil {
				totalAdditionalDamage += additionalDamage.Amount * enemyStats.Hp
			} else {
				totalAdditionalDamage += math.Min(additionalDamage.Amount*enemyStats.Hp, lo.FromPtr(additionalDamage.CappedAmount))
			}
		default:
			panic("invalid additional damage type")
		}
	}
	a.clearAppliedAdditionalDamage()
	return totalAdditionalDamage
}

// clearAppliedAdditionalDamage clears any applied additional damage that only needs to be applied once
func (a AllAdditionalDamage) clearAppliedAdditionalDamage() {
	for additionalDamageName, additionalDamage := range a {
		if additionalDamage.DurationEnd == nil {
			delete(a, additionalDamageName)
		}
	}
}

func (a AllAdditionalDamage) ClearExpiredDurationAdditionalDamage(elapsedTime float64) {
	for additionalDamageName, additionalDamage := range a {
		if additionalDamage.DurationEnd != nil && *additionalDamage.DurationEnd < elapsedTime {
			delete(a, additionalDamageName)
		}
	}
}

type OverTimeDamage struct {
	Damage          float64
	DamageFrequency float64 // time in milliseconds to apply the damage
	DurationStart   float64 // time in milliseconds when the overtime damage should start
	DurationEnd     float64 // time in milliseconds when the overtime damage should end
}

// Result is the result of an attack
type Result struct {
	AttackType       Type
	AttackOption     Option
	DamageDealt      float64
	OvertimeDamage   OverTimeDamage
	AdditionalDamage AdditionalDamage
	Buff             stats.Buff
	Debuffs          []Debuff
}

// CoolDowns is a struct containing the time of a pokemon's attacks
type CoolDowns struct {
	move1CoolDown       float64
	move2CoolDown       float64
	uniteMoveCoolDown   float64
	basicAttackCoolDown float64
}

const BoostedStackDurationBeforeExpiry = 4000.0

type BasicAttack interface {
	Attack(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result Result, err error) // Get the attack dealt by a pokemon's basic attack and possible status effects
}

type SkillMove interface {
	CanCriticallyHit() bool
	IsAvailable(elapsedTime float64) bool                                                                           // Check if the skill move is on cooldown
	Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result Result, err error) // Activate the skill move
}

// GetFramesDelayForAttackSpeed returns the number of frames to wait before attacking (for basic attack) again based on a pokemon's attack speed
func GetFramesDelayForAttackSpeed(attackSpeed float64) int {
	var attackSpeedKey float64
	var foundAttackSpeedKey bool
	for idx, key := range AttackSpeedBucketsKeys {
		if attackSpeed <= key {
			attackSpeedKey = AttackSpeedBucketsKeys[int(math.Max(0, float64(idx-1)))] // Get the prev attack speed bucket
			foundAttackSpeedKey = true
			break
		}
	}
	if !foundAttackSpeedKey {
		return 16 // Max attack speed
	}
	return AttackSpeedBuckets[attackSpeedKey]
}

// AttackSpeedBuckets is a map of attack speed to correlated number of frames to wait before attacking again
// Note: this data is retrieved from this doc: https://docs.google.com/document/d/e/2PACX-1vRM5xkImerqzZoaJhJfMY4dAY3TcsXwtynlvMZhGxXDPVUMMsNNwbhDKiCq1XigZ8zMlE16jierGbnE/pub
var AttackSpeedBuckets = map[float64]int{
	0.0:    60,
	8.1:    56,
	16.42:  52,
	26.11:  48,
	37.56:  44,
	51.29:  40,
	68.05:  36,
	89.04:  32,
	115.99: 28,
	151.81: 24,
	202.04: 20,
	272.51: 16,
}

var AttackSpeedBucketsKeys = []float64{
	0.0,
	8.1,
	16.42,
	26.11,
	37.56,
	51.29,
	68.05,
	89.04,
	115.99,
	151.81,
	202.04,
	272.51,
}
