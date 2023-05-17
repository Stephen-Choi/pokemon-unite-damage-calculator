package pokemon

// StatName is the name of a pokemon stat
type StatName string

const (
	hp                StatName = "hp"
	attack            StatName = "attack"
	defense           StatName = "defense"
	specialAttack     StatName = "specialAttack"
	specialDefense    StatName = "specialDefense"
	attackSpeed       StatName = "attackSpeed"
	criticalHitChance StatName = "criticalHitChance"
	criticalHitDamage StatName = "criticalHitDamage"
	cooldownReduction StatName = "cooldownReduction"
)

// BuffName is the name of the origin of the stat buff applied on a pokemon
type BuffName string

const (
	move1Buff       BuffName = "move1Buff"
	move2Buff       BuffName = "move2Buff"
	uniteMoveBuff   BuffName = "uniteMoveBuff"
	basicAttackBuff BuffName = "basicAttackBuff"
)

// Buffs is a map of buffs applied on a pokemon
type Buffs map[BuffName]map[StatName]float64
