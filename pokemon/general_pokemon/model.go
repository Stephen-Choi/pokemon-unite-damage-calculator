package general_pokemon

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math"
)

type GeneralPokemon struct {
	Name string
	stats.Stats
	stats.Buffs
	attack.AllAdditionalDamage
	BasicAttack attack.BasicAttack
	Move1       attack.SkillMove
	Move2       attack.SkillMove
	UniteMove   attack.SkillMove
	HeldItems   []helditems.HeldItem
	BattleItem  battleitems.BattleItem
}

func (p *GeneralPokemon) GetName() string {
	return p.Name
}

func (p *GeneralPokemon) GetMovesThatCanCrit() (movesThatCanCrit []attack.Option) {
	// Basic attacks can always crit
	movesThatCanCrit = []attack.Option{
		attack.BasicAttackOption,
	}

	if p.Move1.CanCriticallyHit() {
		movesThatCanCrit = append(movesThatCanCrit, attack.Move1)
	}
	if p.Move2.CanCriticallyHit() {
		movesThatCanCrit = append(movesThatCanCrit, attack.Move2)
	}
	if p.UniteMove.CanCriticallyHit() {
		movesThatCanCrit = append(movesThatCanCrit, attack.UniteMove)
	}
	return
}

func (p *GeneralPokemon) GetAvailableActions(elapsedTime float64) (availableAttacks []attack.Option, isBattleItemAvailable bool, err error) {
	// Basic attacks are always available
	availableAttacks = []attack.Option{
		attack.BasicAttackOption,
	}

	pokemonStats := p.GetStats(elapsedTime)

	// Check if skills are available
	if p.Move1.IsAvailable(pokemonStats, elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move1)
	}
	if p.Move2.IsAvailable(pokemonStats, elapsedTime) {
		availableAttacks = append(availableAttacks, attack.Move2)
	}
	if p.UniteMove.IsAvailable(pokemonStats, elapsedTime) {
		availableAttacks = append(availableAttacks, attack.UniteMove)
	}

	// Check if battle item is available
	if p.BattleItem != nil {
		isBattleItemAvailable = p.BattleItem.IsAvailable(elapsedTime)
	}

	return
}

// ActivateBattleItem attempts to activate the battle item
func (p *GeneralPokemon) ActivateBattleItem(elapsedTime float64) {
	_, battleItemEffect, err := p.BattleItem.Activate(p.GetStats(elapsedTime), elapsedTime)
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
func (p *GeneralPokemon) activateHeldItems(statsBeforeAttack stats.Stats, attackResult attack.Result, elapsedTime float64) (debuffs []attack.Debuff, err error) {
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
func (p *GeneralPokemon) Attack(attackOption attack.Option, enemyPokemon enemy.Pokemon, elapsedTime float64) (finalResult attack.Result, err error) {
	statsBeforeAttack := p.GetStats(elapsedTime)

	var attackResult attack.Result
	switch attackOption {
	case attack.BasicAttackOption:
		attackResult, err = p.BasicAttack.Attack(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.Move1:
		attackResult, err = p.Move1.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.Move2:
		attackResult, err = p.Move2.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
	case attack.UniteMove:
		attackResult, err = p.UniteMove.Activate(statsBeforeAttack, enemyPokemon, elapsedTime)
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

func (p *GeneralPokemon) addBuff(attackOption attack.Option, buff stats.Buff) {
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

func (p *GeneralPokemon) addAdditionalDamage(attackOption attack.Option, additionalDamage attack.AdditionalDamage) {
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

// GetStats returns the pokemon's stats with any buffs applied
func (p *GeneralPokemon) GetStats(elapsedTime float64) stats.Stats {
	p.Stats.RemoveExpiredBuffs(p.Buffs, elapsedTime)
	p.Stats.ApplyBuffs(p.Buffs)
	return p.Stats
}

// GetBuffs returns the buffs currently applied on the pokemon
func (p *GeneralPokemon) GetBuffs(elapsedTime float64) stats.Buffs {
	p.Stats.RemoveExpiredBuffs(p.Buffs, elapsedTime)
	return p.Buffs
}
