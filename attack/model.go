package attack

// AttackOption is an enum for the different types of attacks a pokemon can do
type AttackOption string

const (
	Move1       AttackOption = "move1"
	Move2       AttackOption = "move2"
	UniteMove   AttackOption = "uniteMove"
	BasicAttack AttackOption = "basicAttack"
)

// StatusConditions is an enum for the different types of status conditions a pokemon can inflict
type StatusConditions string

const (
	burned StatusConditions = "burned"
)

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

// CoolDowns is a struct containing the cooldowns of a pokemon's attacks
type CoolDowns struct {
	move1CoolDown       float64
	move2CoolDown       float64
	uniteMoveCoolDown   float64
	basicAttackCoolDown float64
}
