package stats

import (
	"errors"
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
	regiBuffDuration               = 90000
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

func (s *Stats) AddStats(stats Stats) {
	s.Hp += stats.Hp
	s.Attack += stats.Attack
	s.Defense += stats.Defense
	s.SpecialAttack += stats.SpecialAttack
	s.SpecialDefense += stats.SpecialDefense
	s.AttackSpeed += stats.AttackSpeed
	s.CriticalHitChance += stats.CriticalHitChance
	s.CriticalHitDamage += stats.CriticalHitDamage
	s.CooldownReduction += stats.CooldownReduction
	s.EnergyRate += stats.EnergyRate
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
func (s *Stats) applyBuff(buff Buff) {
	switch buff.BuffType {
	case PercentIncrease:
		s.Hp *= 1 + buff.StatIncrease.Hp
		s.Attack *= 1 + buff.StatIncrease.Attack
		s.Defense *= 1 + buff.StatIncrease.Defense
		s.SpecialAttack *= 1 + buff.StatIncrease.SpecialAttack
		s.SpecialDefense *= 1 + buff.StatIncrease.SpecialDefense
		s.AttackSpeed *= 1 + buff.StatIncrease.AttackSpeed
		s.CriticalHitChance *= 1 + buff.StatIncrease.CriticalHitChance
		s.CriticalHitDamage *= 1 + buff.StatIncrease.CriticalHitDamage
		s.CooldownReduction *= 1 + buff.StatIncrease.CooldownReduction
	case FlatIncrease:
		s.Hp += buff.StatIncrease.Hp
		s.Attack += buff.StatIncrease.Attack
		s.Defense += buff.StatIncrease.Defense
		s.SpecialAttack += buff.StatIncrease.SpecialAttack
		s.SpecialDefense += buff.StatIncrease.SpecialDefense
		s.AttackSpeed += buff.StatIncrease.AttackSpeed
		s.CriticalHitChance += buff.StatIncrease.CriticalHitChance
		s.CriticalHitDamage += buff.StatIncrease.CriticalHitDamage
		s.CooldownReduction += buff.StatIncrease.CooldownReduction
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
func (s *Stats) removeBuff(buff Buff) {
	switch buff.BuffType {
	case PercentIncrease:
		s.Hp /= 1 + buff.StatIncrease.Hp
		s.Attack /= 1 + buff.StatIncrease.Attack
		s.Defense /= 1 + buff.StatIncrease.Defense
		s.SpecialAttack /= 1 + buff.StatIncrease.SpecialAttack
		s.SpecialDefense /= 1 + buff.StatIncrease.SpecialDefense
		s.AttackSpeed /= 1 + buff.StatIncrease.AttackSpeed
		s.CriticalHitChance /= 1 + buff.StatIncrease.CriticalHitChance
		s.CriticalHitDamage /= 1 + buff.StatIncrease.CriticalHitDamage
		s.CooldownReduction /= 1 + buff.StatIncrease.CooldownReduction
	case FlatIncrease:
		s.Hp -= buff.StatIncrease.Hp
		s.Attack -= buff.StatIncrease.Attack
		s.Defense -= buff.StatIncrease.Defense
		s.SpecialAttack -= buff.StatIncrease.SpecialAttack
		s.SpecialDefense -= buff.StatIncrease.SpecialDefense
		s.AttackSpeed -= buff.StatIncrease.AttackSpeed
		s.CriticalHitChance -= buff.StatIncrease.CriticalHitChance
		s.CriticalHitDamage -= buff.StatIncrease.CriticalHitDamage
		s.CooldownReduction -= buff.StatIncrease.CooldownReduction
	}
}

type BuffType string

const (
	PercentIncrease BuffType = "percentIncrease"
	FlatIncrease    BuffType = "flatIncrease"
)

type Buff struct {
	StatIncrease Stats    `json:"stats-increase"` // Stats is the stats that the buff will apply
	DurationEnd  float64  `json:"duration-end"`   // DurationEnd is a time in milliseconds which holds the time when the buff ends
	BuffType     BuffType `json:"buff-type"`      // BuffType is the type of buff, either percentIncrease or flatIncrease
	Applied      bool     `json:"applied"`        // Applied is a boolean which holds whether the buff has been applied or not
}

func (s Buff) Exists() bool {
	return s != Buff{}
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

const (
	BasicAttackBuff = "basicAttackBuff"
	BattleItemBuff  = "battleItemBuff"
	HeldItem1Buff   = "heldItem1Buff"
	HeldItem2Buff   = "heldItem2Buff"
	HeldItem3Buff   = "heldItem3Buff"
	TeamBuff        = "teamBuff" // Team buff is a buff applied by a teammate TODO not utilized yet...
	RegisteelBuff   = "registeelBuff"
	RegirockBuff    = "regirockBuff"
	RegiceBuff      = "regiceBuff"
)

// Buffs is a map of buffs applied on a pokemon
type Buffs map[string]Buff

func NewBuffs() Buffs {
	return make(Buffs)
}

func (b *Buffs) Add(buffName string, buff Buff) {
	(*b)[buffName] = buff
}

var ValidBuffs = []string{
	"regirock",
	"registeel",
	"regice",
}

var regirockBuff = Buff{
	BuffType: PercentIncrease,
	StatIncrease: Stats{
		Defense:        0.3,
		SpecialDefense: 0.25,
	},
}

var registeelBuff = Buff{
	BuffType: PercentIncrease,
	StatIncrease: Stats{
		Attack:        0.15,
		SpecialAttack: 0.15,
	},
}

var regiceBuff = Buff{
	BuffType: PercentIncrease,
	// TODO: set up healing buff (not really needed now so omitting)
}

// GetTeamBuffs returns a map of buffs applied on a pokemon from a list of buff names
func GetTeamBuffs(teamBuffsNames []string, elapsedTime float64) (Buffs, error) {
	if !isValidTeamBuffs(teamBuffsNames) {
		return nil, errors.New("invalid team buff")
	}

	teamBuffs := NewBuffs()
	for _, teamBuffName := range teamBuffsNames {
		switch teamBuffName {
		case "regirock":
			buff := regirockBuff
			buff.DurationEnd = elapsedTime + regiBuffDuration
			teamBuffs.Add(RegirockBuff, buff)
		case "registeel":
			buff := registeelBuff
			buff.DurationEnd = elapsedTime + regiBuffDuration
			teamBuffs.Add(RegisteelBuff, buff)
		case "regice":
			buff := registeelBuff
			buff.DurationEnd = elapsedTime + regiBuffDuration
			teamBuffs.Add(RegiceBuff, buff)
		}
	}
	return teamBuffs, nil
}

func isValidTeamBuffs(teamBuffs []string) bool {
	for _, buff := range teamBuffs {
		if !isValidBuff(buff) {
			return false
		}
	}
	return true
}

func isValidBuff(buff string) bool {
	for _, validBuff := range ValidBuffs {
		if buff == validBuff {
			return true
		}
	}
	return false
}
