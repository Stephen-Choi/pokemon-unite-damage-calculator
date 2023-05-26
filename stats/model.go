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
	Hp                float64 `json:"hp"`
	Attack            float64 `json:"attack"`
	Defense           float64 `json:"def"`
	SpecialAttack     float64 `json:"sp. attack"`
	SpecialDefense    float64 `json:"sp. def"`
	AttackSpeed       float64 `json:"atk spd"`
	CriticalHitChance float64 `json:"crit chance"`
	CriticalHitDamage float64 `json:"crit dmg"` // Critical attack is default base 200% for all pokemon
	CooldownReduction float64 `json:"CDR"`
	EnergyRate        float64 `json:"energy rate"`
}

func (s *Stats) ApplyBuffs(buffs Buffs) {
	for buffName, buff := range buffs {
		if !buff.Applied {
			s.applyBuff(buff)

			// Set the buff to applied
			statBuff := buffs[buffName]
			statBuff.Applied = true
			buffs[buffName] = statBuff
		}
	}
}

// applyBuff applies the buff to current stats
func (s *Stats) applyBuff(buff StatBuff) {
	switch buff.BuffType {
	case PercentIncrease:
		s.Hp *= 1 + buff.Hp
		s.Attack *= 1 + buff.Attack
		s.Defense *= 1 + buff.Defense
		s.SpecialAttack *= 1 + buff.SpecialAttack
		s.SpecialDefense *= 1 + buff.SpecialDefense
		s.AttackSpeed *= 1 + buff.AttackSpeed
		s.CriticalHitChance *= 1 + buff.CriticalHitChance
		s.CriticalHitDamage *= 1 + buff.CriticalHitDamage
		s.CooldownReduction *= 1 + buff.CooldownReduction
	case FlatIncrease:
		s.Hp += buff.Hp
		s.Attack += buff.Attack
		s.Defense += buff.Defense
		s.SpecialAttack += buff.SpecialAttack
		s.SpecialDefense += buff.SpecialDefense
		s.AttackSpeed += buff.AttackSpeed
		s.CriticalHitChance += buff.CriticalHitChance
		s.CriticalHitDamage += buff.CriticalHitDamage
		s.CooldownReduction += buff.CooldownReduction
	}
}

// RemoveExpiredBuffs removes buffs that have already expired from current stats
func (s *Stats) RemoveExpiredBuffs(buffs Buffs, elapsedTime float64) {
	for buffName, buff := range buffs {
		// If buff is expired
		if buff.DurationEnd < elapsedTime {
			s.removeBuff(buff)
			delete(buffs, buffName) // remove buff from map of buffs
		}
	}
}

// removeBuff removes the buff from current stats
func (s *Stats) removeBuff(buff StatBuff) {
	switch buff.BuffType {
	case PercentIncrease:
		s.Hp /= 1 + buff.Hp
		s.Attack /= 1 + buff.Attack
		s.Defense /= 1 + buff.Defense
		s.SpecialAttack /= 1 + buff.SpecialAttack
		s.SpecialDefense /= 1 + buff.SpecialDefense
		s.AttackSpeed /= 1 + buff.AttackSpeed
		s.CriticalHitChance /= 1 + buff.CriticalHitChance
		s.CriticalHitDamage /= 1 + buff.CriticalHitDamage
		s.CooldownReduction /= 1 + buff.CooldownReduction
	case FlatIncrease:
		s.Hp -= buff.Hp
		s.Attack -= buff.Attack
		s.Defense -= buff.Defense
		s.SpecialAttack -= buff.SpecialAttack
		s.SpecialDefense -= buff.SpecialDefense
		s.AttackSpeed -= buff.AttackSpeed
		s.CriticalHitChance -= buff.CriticalHitChance
		s.CriticalHitDamage -= buff.CriticalHitDamage
		s.CooldownReduction -= buff.CooldownReduction
	}
}

type BuffType string

const (
	PercentIncrease BuffType = "percentIncrease"
	FlatIncrease    BuffType = "flatIncrease"
)

type StatBuff struct {
	Stats
	DurationEnd float64  // DurationEnd is a time in milliseconds which holds the time when the buff ends
	BuffType    BuffType // BuffType is the type of buff, either percentIncrease or flatIncrease
	Applied     bool
}

func (s StatBuff) Exists() bool {
	return s != StatBuff{}
}

// ToTypedStats converts the JsonStats struct to the internal use typed Stats struct
func ToTypedStats(jsonStats JsonStats) (Stats, error) {
	level, err := strconv.Atoi(jsonStats.Level)
	if err != nil {
		fmt.Println("Error parsing level:", err)
		return Stats{}, err
	}
	hp, err := strconv.ParseFloat(jsonStats.Hp, 64)
	if err != nil {
		fmt.Println("Error parsing hp:", err)
		return Stats{}, err
	}
	attack, err := strconv.ParseFloat(jsonStats.Attack, 64)
	if err != nil {
		fmt.Println("Error parsing attack:", err)
		return Stats{}, err
	}
	defense, err := strconv.ParseFloat(jsonStats.Defense, 64)
	if err != nil {
		fmt.Println("Error parsing defense:", err)
		return Stats{}, err
	}
	specialAttack, err := strconv.ParseFloat(jsonStats.SpecialAttack, 64)
	if err != nil {
		fmt.Println("Error parsing special attack:", err)
		return Stats{}, err
	}
	specialDefense, err := strconv.ParseFloat(jsonStats.SpecialDefense, 64)
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
	Move1Buff       BuffName = "move1Buff"
	Move2Buff       BuffName = "move2Buff"
	UniteMoveBuff   BuffName = "uniteMoveBuff"
	BasicAttackBuff BuffName = "basicAttackBuff"
)

// Buffs is a map of buffs applied on a pokemon
type Buffs map[BuffName]StatBuff

func (b Buffs) AddBuff(buffName BuffName, buff StatBuff) {
	b[buffName] = buff
}
