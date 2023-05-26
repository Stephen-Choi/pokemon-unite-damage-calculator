package attack

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
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

// DebuffType is an enum for the different types of status conditions a pokemon can inflict
type DebuffType string

const (
	IgnoreDefense DebuffType = "ignoreDefense"
)

// Debuff is a struct containing the type and amount of debuff to be applied
type Debuff struct {
	DebuffType DebuffType
	stats.Stats
}

type AdditionalDamageType string

const (
	SimpleAdditionalDamage AdditionalDamageType = "single"             // SingleAdditionalDamage is a single instance of additional damage
	PercentDamageBoost     AdditionalDamageType = "percentDamageBoost" // PercentDamageBoost is a damage boost (as a percent increase) that lasts for a certain amount of time
	RemainingEnemyHp       AdditionalDamageType = "remainingEnemyHp"   // RemainingEnemyHp is additional damage that scales with the enemy's remaining HP
)

// AdditionalDamage is a struct containing the type and amount of additional damage to be applied
type AdditionalDamage struct {
	Type         AdditionalDamageType
	Amount       float64
	CappedAmount *float64 // only applicable to certain held items (i.e muscle band)
	Duration     *float64 // only applicable to certain held items (i.e energy amplifier)
}

type OverTimeDamage struct {
	Damage          float64
	DamageFrequency float64 // time in milliseconds to apply the damage
	DurationEnd     float64 // time in milliseconds when the overtime damage should end
}

// Result is the result of an attack
type Result struct {
	AttackType     Type
	AttackOption   Option
	DamageDealt    float64
	OvertimeDamage OverTimeDamage
	Buff           stats.Buff
	Debuffs        []Debuff
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
	IsAvailable(elapsedTime float64) bool                                                                           // Check if the skill move is on cooldown
	Activate(originalStats stats.Stats, enemyPokemon enemy.Pokemon, elapsedTime float64) (result Result, err error) // Activate the skill move
}
