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
	Passive     attack.Passive
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

	pokemonStats := p.GetStats()

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
	_, battleItemEffect, err := p.BattleItem.Activate(p.GetStats(), elapsedTime)
	if err != nil {
		return
	}

	if battleItemEffect.Buff.Exists() {
		p.Buffs.Add(p.BattleItem.GetName(), battleItemEffect.Buff)
	}
	if battleItemEffect.AdditionalDamage.Exists() {
		p.AllAdditionalDamage.Add(p.BattleItem.GetName(), battleItemEffect.AdditionalDamage)
	}
}

// activateHeldItems attempts to activate held items
// Note: since most held items are triggered automatically, no need to export this method. Can simply call it after an attack occurs
func (p *GeneralPokemon) activateHeldItems(statsBeforeAttack stats.Stats, attackResult attack.Result, elapsedTime float64) (debuffs []attack.Debuff, err error) {
	attackOption := attackResult.AttackOption
	attackType := attackResult.AttackType

	// Default to the attack result, build on top of it
	for _, heldItem := range p.HeldItems {
		_, heldItemEffect, err := heldItem.Activate(statsBeforeAttack, elapsedTime, attackOption, attackType, attackResult.BaseDamageDealt)
		if err != nil {
			return nil, err
		}

		// Apply any buffs that were granted from the held item
		if heldItemEffect.Buff.Exists() {
			p.Buffs.Add(heldItem.GetName(), heldItemEffect.Buff)
		}

		// Apply any additional damage that were granted from the held item
		if heldItemEffect.AdditionalDamage.Exists() {
			p.AllAdditionalDamage.Add(heldItem.GetName(), heldItemEffect.AdditionalDamage)
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
	statsBeforeAttack := p.GetStats()
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
	if attackResult.AdditionalDamageEffect.Exists() {
		p.addAdditionalDamage(attackOption, attackResult.AdditionalDamageEffect)
	}

	// Apply passive effects if available
	if p.Passive.IsAvailable(elapsedTime) {
		passiveResult, err := p.Passive.Activate(statsBeforeAttack, attackResult, elapsedTime)
		if err != nil {
			return finalResult, err
		}
		if passiveResult.Buff.Exists() {
			p.addBuff(attackOption, passiveResult.Buff)
		}
		if passiveResult.AdditionalDamageEffect.Exists() {
			p.addAdditionalDamage(attackOption, passiveResult.AdditionalDamageEffect)
		}
		for _, debuff := range passiveResult.Debuffs {
			attackResult.Debuffs = append(attackResult.Debuffs, debuff)
		}
	}

	// Apply held item effects
	heldItemDebuffs, err := p.activateHeldItems(statsBeforeAttack, attackResult, elapsedTime)

	// Set up the final result of the attack with all debuffs and additional damage applied
	finalResult = attackResult

	for _, debuff := range heldItemDebuffs {
		finalResult.Debuffs = append(finalResult.Debuffs, debuff)
	}

	// Apply additional damage
	totalAdditionalDamage := p.AllAdditionalDamage.Calculate(attackResult.BaseDamageDealt, attackResult.AttackOption, enemyPokemon.GetStats(), enemyPokemon.IsWild())
	finalResult.AdditionalDamageDealt = totalAdditionalDamage

	// Round to 2 decimal places
	finalResult.BaseDamageDealt = math.Floor(finalResult.BaseDamageDealt*100) / 100
	finalResult.AdditionalDamageDealt = math.Floor(finalResult.AdditionalDamageDealt*100) / 100

	return
}

func (p *GeneralPokemon) addBuff(attackOption attack.Option, buff stats.Buff) {
	switch attackOption {
	case attack.BasicAttackOption:
		p.Buffs.Add(stats.BasicAttackBuff, buff)
	case attack.Move1:
		p.Buffs.Add(p.Move1.GetName(), buff)
	case attack.Move2:
		p.Buffs.Add(p.Move2.GetName(), buff)
	case attack.UniteMove:
		p.Buffs.Add(p.UniteMove.GetName(), buff)
	}
}

func (p *GeneralPokemon) addAdditionalDamage(attackOption attack.Option, additionalDamage attack.AdditionalDamage) {
	switch attackOption {
	case attack.BasicAttackOption:
		p.AllAdditionalDamage.Add(string(attack.BasicAttackOption), additionalDamage)
	case attack.Move1:
		p.AllAdditionalDamage.Add(p.Move1.GetName(), additionalDamage)
	case attack.Move2:
		p.AllAdditionalDamage.Add(p.Move2.GetName(), additionalDamage)
	case attack.UniteMove:
		p.AllAdditionalDamage.Add(p.UniteMove.GetName(), additionalDamage)
	}
}

// GetStats returns the pokemon's stats with any buffs applied
func (p *GeneralPokemon) GetStats() stats.Stats {
	p.Stats.ApplyBuffs(p.Buffs)
	return p.Stats
}

// GetBuffs returns the buffs currently applied on the pokemon
func (p *GeneralPokemon) GetBuffs() stats.Buffs {
	return p.Buffs
}

// GetAllAdditionalDamage returns the additional damage currently applied on the pokemon
func (p *GeneralPokemon) GetAllAdditionalDamage() attack.AllAdditionalDamage {
	return p.AllAdditionalDamage
}

// ClearExpiredEffects clears expired buffs and additional damage effects
func (p *GeneralPokemon) ClearExpiredEffects(elapsedTime float64) {
	p.Stats.RemoveExpiredBuffs(p.Buffs, elapsedTime)
	p.AllAdditionalDamage.ClearCompletedAdditionalDamageEffects(elapsedTime)
}

// AddAdditionalDamage clears expired buffs and additional damage effects
func (p *GeneralPokemon) AddAdditionalDamage(additionalDamageName string, additionalDamage attack.AdditionalDamage) {
	p.AllAdditionalDamage.Add(additionalDamageName, additionalDamage)
}
