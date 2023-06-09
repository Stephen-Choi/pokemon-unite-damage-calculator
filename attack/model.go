package attack

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/time"
	"github.com/samber/lo"
	"math"
)

const (
	BasicAttackName = "basic attack"
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
	SingleInstance                                    AdditionalDamageType = "singleInstance"                                    // SingleInstance is a single instance of additional damage
	PercentDamageBoost                                AdditionalDamageType = "percentDamageBoost"                                // PercentDamageBoost is a damage boost (as a percent increase) that lasts for a certain amount of time
	RemainingEnemyHp                                  AdditionalDamageType = "remainingEnemyHp"                                  // RemainingEnemyHp is additional damage that scales with the enemy's remaining HP
	DefenderAndSupporterDamageBoostAgainstWildPokemon AdditionalDamageType = "defenderAndSupporterDamageBoostAgainstWildPokemon" // DefenderAndSupporterDamageBoost is a damage boost that applies to defenders and supporters on the team against wild pokemon
)

func GetDefenderAndSupporterWildDamageBoost(level int) AdditionalDamage {
	var damagePercentBoost float64
	if level >= 9 {
		damagePercentBoost = 0.2
	} else {
		damagePercentBoost = 0.1
	}

	return AdditionalDamage{
		Type:        DefenderAndSupporterDamageBoostAgainstWildPokemon,
		Amount:      damagePercentBoost,
		DurationEnd: lo.ToPtr(float64(time.FullGameTimeInMilliseconds)),
	}
}

// AdditionalDamage is a struct containing the type and amount of additional damage to be applied
type AdditionalDamage struct {
	Type         AdditionalDamageType `json:"type"`
	Amount       float64              `json:"amount"`
	CappedAmount *float64             `json:"capped-amount,omitempty"` // only applicable to certain held items (i.e muscle band)
	DurationEnd  *float64             `json:"duration-end,omitempty"`  // only applicable to certain held items (i.e energy amplifier)
}

func (a AdditionalDamage) Exists() bool {
	return a != AdditionalDamage{}
}

type AllAdditionalDamage map[string]AdditionalDamage

func NewAllAdditionalDamage() AllAdditionalDamage {
	return make(AllAdditionalDamage)
}

func (a AllAdditionalDamage) Add(additionalDamageName string, additionalDamage AdditionalDamage) {
	a[additionalDamageName] = additionalDamage
}

// Calculate calculates the total additional damage
func (a AllAdditionalDamage) Calculate(baseAttackDamage float64, attackOption Option, enemyStats stats.Stats, enemyIsWildPokemon bool) float64 {
	totalAdditionalDamage := 0.0
	for _, additionalDamage := range a {
		switch additionalDamage.Type {
		case SingleInstance:
			totalAdditionalDamage += additionalDamage.Amount
		case PercentDamageBoost:
			totalAdditionalDamage += baseAttackDamage * additionalDamage.Amount
		case RemainingEnemyHp:
			if additionalDamage.CappedAmount == nil {
				totalAdditionalDamage += additionalDamage.Amount * enemyStats.Hp
			} else {
				totalAdditionalDamage += math.Min(additionalDamage.Amount*enemyStats.Hp, lo.FromPtr(additionalDamage.CappedAmount))
			}
		case DefenderAndSupporterDamageBoostAgainstWildPokemon:
			if enemyIsWildPokemon && attackOption == Move1 || attackOption == BasicAttackOption {
				totalAdditionalDamage += baseAttackDamage * additionalDamage.Amount
			}
		default:
			panic("invalid additional damage type")
		}
	}
	return totalAdditionalDamage
}

func (a AllAdditionalDamage) clearSingleInstanceAppliedAdditionalDamage() {
	for additionalDamageName, additionalDamage := range a {
		if additionalDamage.DurationEnd == nil {
			delete(a, additionalDamageName)
		}
	}
}

func (a AllAdditionalDamage) clearExpiredDurationAdditionalDamage(elapsedTime float64) {
	for additionalDamageName, additionalDamage := range a {
		if additionalDamage.DurationEnd != nil && *additionalDamage.DurationEnd <= elapsedTime {
			delete(a, additionalDamageName)
		}
	}
}

// ClearCompletedAdditionalDamageEffects clears any additional damage effects that have been completed
func (a AllAdditionalDamage) ClearCompletedAdditionalDamageEffects(elapsedTime float64) {
	a.clearExpiredDurationAdditionalDamage(elapsedTime)
	a.clearSingleInstanceAppliedAdditionalDamage()
}

type OverTimeDamage struct {
	Source                  string  `json:"source"`
	AttackType              Type    `json:"attack-type"`
	BaseDamage              float64 `json:"base-damage"`
	LastInflictedDamageTime float64 `json:"last-inflicted-damage-time"`
	DamageFrequency         float64 `json:"damage-frequency"` // time in milliseconds to apply the damage
	DurationStart           float64 `json:"duration-start"`   // time in milliseconds when the overtime damage should start
	DurationEnd             float64 `json:"duration-end"`     // time in milliseconds when the overtime damage should end
}

func (o OverTimeDamage) Exists() bool {
	return o != OverTimeDamage{}
}

type ExecutePercentDamage struct {
	Percent      float64
	CappedDamage float64
}

func (e ExecutePercentDamage) Exists() bool {
	return e != ExecutePercentDamage{}
}

// Result is the result of an attack
type Result struct {
	AttackType             Type
	AttackOption           Option
	AttackName             string
	BaseDamageDealt        float64
	AdditionalDamageDealt  float64
	OvertimeDamage         OverTimeDamage
	AdditionalDamageEffect AdditionalDamage
	Buff                   stats.Buff
	Debuffs                []Debuff
	AttackDuration         float64 // time in milliseconds that the attack took to complete TODO: set this for all moves if ever someone gets the frame data for each move
	ExecutionPercentDamage ExecutePercentDamage
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
	// TODO: Add a setBoostedStacks method since some skills lead to boosted autos
	Attack(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result Result, err error) // Get the attack dealt by a pokemon's basic attack and possible status effects
}

type SkillMove interface {
	GetName() string
	CanCriticallyHit() bool
	IsAvailable(pokemonStats stats.Stats, elapsedTime float64) bool                                                // Check if the skill move is on cooldown
	Activate(pokemonStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result Result, err error) // Activate the skill move
}

// GetDelayForAttackSpeed returns the time delay to wait before attacking (for basic attack) again based on a pokemon's attack speed
func GetDelayForAttackSpeed(attackSpeed float64) float64 {
	var attackSpeedKey float64
	var foundAttackSpeedKey bool
	for idx, key := range AttackSpeedBucketsKeys {
		if attackSpeed <= key {
			attackSpeedKey = AttackSpeedBucketsKeys[int(math.Max(0, float64(idx-1)))] // Get the prev attack speed bucket
			foundAttackSpeedKey = true
			break
		}
	}
	var frameDelay int
	if !foundAttackSpeedKey {
		frameDelay = 16 // Max attack speed
	} else {
		frameDelay = AttackSpeedBuckets[attackSpeedKey]
	}

	// Every 4 frames occurs in 66.67 milliseconds
	return float64(frameDelay) / 4 * 66.67
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

// ApplyDebuffs applies debuffs to the enemy pokemon for a specific attacking pokemon
// This needs to occur on a per pokemon basis since a pokemon can have debuff effects that only apply for itself (ex. slick spoon)
func ApplyDebuffs(attackingPokemon string, enemy enemy.Pokemon, debuffs []Debuff) (enemyStats stats.Stats) {
	enemyStats = enemy.GetStats()

	// Debuffs apply additively
	debuffStats := stats.Stats{}

	for _, debuff := range debuffs {
		switch debuff.DebuffEffect {
		case IgnoreDefenseForAttackingPokemon:
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
