package pokemon

// StatName is the name of a pokemon stat
type StatName string

const (
	hpStat                StatName = "hp"
	attackStat            StatName = "attack"
	defenseStat           StatName = "defense"
	specialAttackStat     StatName = "specialAttack"
	specialDefenseStat    StatName = "specialDefense"
	attackSpeedStat       StatName = "attackSpeed"
	criticalHitChanceStat StatName = "criticalHitChance"
	criticalHitDamageStat StatName = "criticalHitDamage"
	cooldownReductionStat StatName = "cooldownReduction"
)

// Stats is a struct containing the stats of a pokemon
type Stats struct {
	level             int     `json:"level"`
	hp                int     `json:"hp"`
	attack            int     `json:"attack"`
	defense           int     `json:"def"`
	specialAttack     int     `json:"sp. attack"`
	specialDefense    int     `json:"sp. def"`
	attackSpeed       float64 `json:"atk spd"`
	criticalHitChance float64 `json:"crit chance"`
	criticalHitDamage float64 `json:"-"` // Critical damage is default base 200% for all pokemon
	cooldownReduction float64 `json:"CDR"`
}

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
