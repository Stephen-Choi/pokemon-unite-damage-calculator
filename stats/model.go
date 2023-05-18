package stats

import (
	"fmt"
	"strconv"
)

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

// JsonStats is a temporary struct used to unmarshal the pokemon json stats
type JsonStats struct {
	Level             string `json:"level"`
	Hp                string `json:"hp"`
	Attack            string `json:"attack"`
	Defense           string `json:"def"`
	SpecialAttack     string `json:"sp. attack"`
	SpecialDefense    string `json:"sp. def"`
	AttackSpeed       string `json:"atk spd"`
	CriticalHitChance string `json:"crit chance"`
	CriticalHitDamage string `json:"crit dmg"`
	CooldownReduction string `json:"CDR"`
}

// Stats is a struct containing the stats of a pokemon
type Stats struct {
	Level             int     `json:"level"`
	Hp                int     `json:"hp"`
	Attack            int     `json:"attack"`
	Defense           int     `json:"def"`
	SpecialAttack     int     `json:"sp. attack"`
	SpecialDefense    int     `json:"sp. def"`
	AttackSpeed       float64 `json:"atk spd"`
	CriticalHitChance float64 `json:"crit chance"`
	CriticalHitDamage float64 `json:"crit dmg"` // Critical attack is default base 200% for all pokemon
	CooldownReduction float64 `json:"CDR"`
	EnergyRate        float64 `json:"energy rate"`
}

// ToTypedStats converts the JsonStats struct to the internal use typed Stats struct
func ToTypedStats(jsonStats JsonStats) (Stats, error) {
	level, err := strconv.Atoi(jsonStats.Level)
	if err != nil {
		fmt.Println("Error parsing level:", err)
		return Stats{}, err
	}
	hp, err := strconv.Atoi(jsonStats.Hp)
	if err != nil {
		fmt.Println("Error parsing hp:", err)
		return Stats{}, err
	}
	attack, err := strconv.Atoi(jsonStats.Attack)
	if err != nil {
		fmt.Println("Error parsing attack:", err)
		return Stats{}, err
	}
	defense, err := strconv.Atoi(jsonStats.Defense)
	if err != nil {
		fmt.Println("Error parsing defense:", err)
		return Stats{}, err
	}
	specialAttack, err := strconv.Atoi(jsonStats.SpecialAttack)
	if err != nil {
		fmt.Println("Error parsing special attack:", err)
		return Stats{}, err
	}
	specialDefense, err := strconv.Atoi(jsonStats.SpecialDefense)
	if err != nil {
		fmt.Println("Error parsing special defense:", err)
		return Stats{}, err
	}
	attackSpeed, err := convertPercentToFloat(jsonStats.AttackSpeed)
	if err != nil {
		fmt.Println("Error parsing attack speed:", err)
		return Stats{}, err
	}
	criticalHitChance, err := convertPercentToFloat(jsonStats.CriticalHitChance)
	if err != nil {
		fmt.Println("Error parsing critical hit chance:", err)
		return Stats{}, err
	}
	cooldownReduction, err := convertPercentToFloat(jsonStats.CooldownReduction)
	if err != nil {
		fmt.Println("Error parsing cooldown reduction:", err)
		return Stats{}, err
	}

	return Stats{
		Level:             level,
		Hp:                hp,
		Attack:            attack,
		Defense:           defense,
		SpecialAttack:     specialAttack,
		SpecialDefense:    specialDefense,
		AttackSpeed:       attackSpeed,
		CriticalHitChance: criticalHitChance,
		CriticalHitDamage: 2.0,
		CooldownReduction: cooldownReduction,
	}, nil
}

func convertPercentToFloat(percentString string) (float64, error) {
	percentString = percentString[:len(percentString)-1] // Remove the '%' symbol
	percentValue, err := strconv.ParseFloat(percentString, 64)
	if err != nil {
		fmt.Println("Error parsing percent:", err)
		return 0, err
	}

	// Convert percent to a decimal value
	decimalValue := percentValue / 100.0
	return decimalValue, nil
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
