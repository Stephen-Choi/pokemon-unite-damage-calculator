package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	stats "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

type MoveName string

const (
	ThunderShockName MoveName = "thunder_shock"
	ElectroBallName  MoveName = "electro_ball"
	ThunderName      MoveName = "thunder"
	ElectrowebName   MoveName = "electroweb"
	VoltTackleName   MoveName = "volt_tackle"
	ThunderboltName  MoveName = "thunderbolt"
)

var move1Set = []MoveName{
	ThunderShockName,
	ElectroBallName,
	ThunderName,
}

var move2Set = []MoveName{
	ElectrowebName,
	VoltTackleName,
	ThunderboltName,
}

// Pikachu is a pokemon
type Pikachu struct {
	stats.Stats
	stats.Buffs
	basicAttack attack.BasicAttack
	move1       attack.SkillMove
	move2       attack.SkillMove
	uniteMove   attack.SkillMove
	HeldItems   []helditems.HeldItem
	BattleItem  battleitems.BattleItem
}

func NewPikachu(level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem, emblems stats.Stats) (p *Pikachu, err error) {
	// Get pokemon stats
	pokemonStats, err := pokemon.FetchPokemonStats(pokemon.PikachuName, level)
	if err != nil {
		return
	}

	// TODO: apply emblems

	// Get move 1
	if !move1Exists(move1Name) {
		err = errors.New("invalid move 1")
		return
	}
	move1, err := getMove(move1Name, level)
	if err != nil {
		return
	}

	// Get move 2
	if !move2Exists(move2Name) {
		err = errors.New("invalid move 2")
		return
	}
	move2, err := getMove(move2Name, level)
	if err != nil {
		return
	}

	p = &Pikachu{
		Stats:       pokemonStats,
		basicAttack: NewBasicAttack(),
		move1:       move1,
		move2:       move2,
		uniteMove:   NewThunderstorm(level),
		HeldItems:   heldItems,
		BattleItem:  battleItem,
	}

	return
}

func move1Exists(move string) bool {
	typedMove := MoveName(move)
	for _, viableMove1 := range move1Set {
		if typedMove == viableMove1 {
			return true
		}
	}
	return false
}

func move2Exists(move string) bool {
	typedMove := MoveName(move)
	for _, viableMove2 := range move2Set {
		if typedMove == viableMove2 {
			return true
		}
	}
	return false
}

func getMove(moveName string, level int) (move attack.SkillMove, err error) {
	typedMoveName := MoveName(moveName)

	switch typedMoveName {
	case ThunderShockName:
		move, err = NewThunderShock(level)
	case ElectroBallName:
		move, err = NewElectroBall(level)
	case ThunderName:
		move, err = NewThunder(level)
	case ElectrowebName:
		move, err = NewElectroweb(level)
	case VoltTackleName:
		move, err = NewVoltTackle(level)
	case ThunderboltName:
		move, err = NewThunderbolt(level)
	default:
		err = errors.New("invalid move name")
	}
	return
}

func (p *Pikachu) GetAvailableAttacks(elapsedTime float64) (availableAttacks []attack.Option, err error) {
	// Basic attacks are always available
	availableAttacks = []attack.Option{
		attack.BasicAttackOption,
	}

	if p.move1.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move1)
	}
	if p.move2.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move2)
	}
	if p.uniteMove.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.UniteMove)
	}

	return
}

func (p *Pikachu) Attack(attackOption attack.Option, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error) {
	switch attackOption {
	case attack.BasicAttackOption:
		result, err = p.basicAttack.Attack(p.getStats(), enemyPokemon, elapsedTime)
	case attack.Move1:
		result, err = p.move1.Activate(p.getStats(), enemyPokemon, elapsedTime)
	case attack.Move2:
		result, err = p.move2.Activate(p.getStats(), enemyPokemon, elapsedTime)
	case attack.UniteMove:
		result, err = p.uniteMove.Activate(p.getStats(), enemyPokemon, elapsedTime)
	default:
		err = errors.New("invalid attack option")
	}

	if err != nil {
		return
	}

	if result.Buff.Exists() {
		p.addBuff(attackOption, result.Buff)
	}

	return
}

func (p *Pikachu) addBuff(attackOption attack.Option, buff stats.StatBuff) {
	switch attackOption {
	case attack.BasicAttackOption:
		p.Buffs.AddBuff(stats.BasicAttackBuff, buff)
	case attack.Move1:
		p.Buffs.AddBuff(stats.Move1Buff, buff)
	case attack.Move2:
		p.Buffs.AddBuff(stats.Move2Buff, buff)
	case attack.UniteMove:
		p.Buffs.AddBuff(stats.UniteMoveBuff, buff)
	}
}

// getStats returns the pokemon's stats with any buffs applied
func (p *Pikachu) getStats() stats.Stats {
	p.Stats.ApplyBuffs(p.Buffs)
	return p.Stats
}
