package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	stats "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math"
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
	Name string
	stats.Stats
	stats.Buffs
	attack.AllAdditionalDamage
	basicAttack attack.BasicAttack
	move1       attack.SkillMove
	move2       attack.SkillMove
	uniteMove   attack.SkillMove
	HeldItems   []helditems.HeldItem
	BattleItem  battleitems.BattleItem
}

func NewPikachu(level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem, emblems *stats.Stats) (p *Pikachu, err error) {
	// Get pokemon stats
	pokemonStats, err := pokemon.FetchPokemonStats(pokemon.PikachuName, level)
	if err != nil {
		return
	}

	// Apply held item stats
	for _, heldItem := range heldItems {
		pokemonStats.AddStats(heldItem.GetStatBoosts(pokemonStats))
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
		Name:                pokemon.PikachuName,
		Stats:               pokemonStats,
		basicAttack:         NewBasicAttack(),
		move1:               move1,
		move2:               move2,
		uniteMove:           NewThunderstorm(level),
		HeldItems:           heldItems,
		BattleItem:          battleItem,
		Buffs:               stats.NewBuffs(),
		AllAdditionalDamage: attack.NewAllAdditionalDamage(),
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

func (p *Pikachu) GetAvailableActions(elapsedTime float64) (availableAttacks []attack.Option, isBattleItemAvailable bool, err error) {
	// Basic attacks are always available
	availableAttacks = []attack.Option{
		attack.BasicAttackOption,
	}

	// Check if skills are available
	if p.move1.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move1)
	}
	if p.move2.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move2)
	}
	if p.uniteMove.IsAvailable(elapsedTime) {
		availableAttacks = append(availableAttacks, attack.UniteMove)
	}

	// Check if battle item is available
	isBattleItemAvailable = p.BattleItem.IsAvailable(elapsedTime)

	return
}

// ActivateBattleItem attempts to activate the battle item
func (p *Pikachu) ActivateBattleItem(elapsedTime float64) {
	_, battleItemEffect, err := p.BattleItem.Activate(p.getStats(elapsedTime), elapsedTime)
	if err != nil {
		return
	}

	if battleItemEffect.Buff.Exists() {
		p.Buffs.Add(stats.BattleItemBuff, battleItemEffect.Buff)
	}
	if battleItemEffect.AdditionalDamage.Exists() {
		p.AllAdditionalDamage.Add(attack.BattleItemAdditionalDamage, battleItemEffect.AdditionalDamage)
	}
}

// activateHeldItems attempts to activate held items
// Note: since most held items are triggered automatically, no need to export this method. Can simply call it after an attack occurs
func (p *Pikachu) activateHeldItems(statsBeforeAttack stats.Stats, attackResult attack.Result, elapsedTime float64) (debuffs []attack.Debuff, err error) {
	attackOption := attackResult.AttackOption
	attackType := attackResult.AttackType

	// Default to the attack result, build on top of it
	for index, heldItem := range p.HeldItems {
		_, heldItemEffect, err := heldItem.Activate(statsBeforeAttack, elapsedTime, attackOption, attackType)
		if err != nil {
			return nil, err
		}

		// Apply any buffs that were granted from the held item
		if heldItemEffect.Buff.Exists() {
			p.Buffs.Add(stats.GetHeldItemName(index), heldItemEffect.Buff)
		}

		// Apply any additional damage that were granted from the held item
		if heldItemEffect.AdditionalDamage.Exists() {
			p.AllAdditionalDamage.Add(attack.GetAdditionalDamageName(index), heldItemEffect.AdditionalDamage)
		}

		// Add debuffs that are inflicted by held item to the enemy pokemon
		if heldItemEffect.Debuff.Exists() {
			debuff := heldItemEffect.Debuff
			debuff.FromPokemon = p.Name
			debuffs = append(debuffs, debuff)
		}
	}
	return
}

// Attack performs an attack on an opposing pokemon
// Note: Any buffs that are granted from the attack should not apply in this iteration of the attack. They should only
// be available for the next attack (thus why we have a statsBeforeAttack var that is used everywhere that stats are needed).
func (p *Pikachu) Attack(attackOption attack.Option, enemyPokemon enemy.Pokemon, elapsedTime float64) (finalResult attack.Result, err error) {
	statsBeforeAttack := p.getStats(elapsedTime)

	var attackResult attack.Result
	switch attackOption {
	case attack.BasicAttackOption:
		attackResult, err = p.basicAttack.Attack(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.Move1:
		attackResult, err = p.move1.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.Move2:
		attackResult, err = p.move2.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.UniteMove:
		attackResult, err = p.uniteMove.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
	default:
		err = errors.New("invalid attack option")
	}
	if err != nil {
		return
	}

	// Add attack related buffs
	if attackResult.Buff.Exists() {
		p.addBuff(attackOption, attackResult.Buff)
	}
	// Add additional damage
	if attackResult.AdditionalDamage.Exists() {
		p.addAdditionalDamage(attackOption, attackResult.AdditionalDamage)
	}

	// Apply held item effects
	heldItemDebuffs, err := p.activateHeldItems(statsBeforeAttack, attackResult, elapsedTime)

	// Set up the final result of the attack with all debuffs and additional damage applied
	finalResult = attackResult

	for _, debuff := range heldItemDebuffs {
		finalResult.Debuffs = append(finalResult.Debuffs, debuff)
	}

	// Apply additional damage
	p.AllAdditionalDamage.ClearExpiredDurationAdditionalDamage(elapsedTime)
	totalAdditionalDamage := p.AllAdditionalDamage.Calculate(attackResult.DamageDealt, enemyPokemon.GetStats())
	finalResult.DamageDealt += totalAdditionalDamage

	// Round to 2 decimal places
	finalResult.DamageDealt = math.Floor(finalResult.DamageDealt*100) / 100

	return
}

func (p *Pikachu) addBuff(attackOption attack.Option, buff stats.Buff) {
	switch attackOption {
	case attack.BasicAttackOption:
		p.Buffs.Add(stats.BasicAttackBuff, buff)
	case attack.Move1:
		p.Buffs.Add(stats.Move1Buff, buff)
	case attack.Move2:
		p.Buffs.Add(stats.Move2Buff, buff)
	case attack.UniteMove:
		p.Buffs.Add(stats.UniteMoveBuff, buff)
	}
}

func (p *Pikachu) addAdditionalDamage(attackOption attack.Option, additionalDamage attack.AdditionalDamage) {
	switch attackOption {
	case attack.BasicAttackOption:
		p.AllAdditionalDamage.Add(attack.BasicAttackAdditionalDamage, additionalDamage)
	case attack.Move1:
		p.AllAdditionalDamage.Add(attack.Move1AdditionalDamage, additionalDamage)
	case attack.Move2:
		p.AllAdditionalDamage.Add(attack.Move2AdditionalDamage, additionalDamage)
	case attack.UniteMove:
		p.AllAdditionalDamage.Add(attack.UniteMoveAdditionalDamage, additionalDamage)
	}
}

// getStats returns the pokemon's stats with any buffs applied
func (p *Pikachu) getStats(elapsedTime float64) stats.Stats {
	p.Stats.RemoveExpiredBuffs(p.Buffs, elapsedTime)
	p.Stats.ApplyBuffs(p.Buffs)
	return p.Stats
}
