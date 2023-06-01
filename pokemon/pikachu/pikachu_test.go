package pikachu

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setUpPikachu(t *testing.T, level int) *Pikachu {
	// Set up held items
	muslceBand, err := helditems.NewMuscleBand()
	assert.NoError(t, err)
	energyAmp, err := helditems.NewEnergyAmplifier()
	assert.NoError(t, err)
	expShare, err := helditems.NewExpShare()
	assert.NoError(t, err)

	// Set up battle item
	xAttack, err := battleitems.NewXAttack()
	assert.NoError(t, err)

	pikachu, err := NewPikachu(level, string(ElectroBallName), string(VoltTackleName), []helditems.HeldItem{muslceBand, energyAmp, expShare}, xAttack, nil)
	assert.NoError(t, err)

	return pikachu
}

func Test_NewPikachu(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		level := 9
		pikachu := setUpPikachu(t, level)

		// Assert stats are as expected
		assert.Equal(t, stats.Stats{
			Level:             level,
			Hp:                4756,
			Attack:            213,
			Defense:           92,
			SpecialAttack:     418,
			SpecialDefense:    72,
			AttackSpeed:       0.175,
			CriticalHitChance: 0,
			CriticalHitDamage: 2.0,
			CooldownReduction: 0.195,
			EnergyRate:        0.06,
		}, pikachu.Stats)

		// Assert items are as expected
		muslceBand, _ := helditems.NewMuscleBand()
		energyAmp, _ := helditems.NewEnergyAmplifier()
		expShare, _ := helditems.NewExpShare()
		assert.Equal(t, []helditems.HeldItem{muslceBand, energyAmp, expShare}, pikachu.HeldItems)

		// Assert moves are as expected
		electroBall, err := NewElectroBall(level)
		assert.NoError(t, err)
		assert.Equal(t, electroBall, pikachu.Move1)

		voltTackle, err := NewVoltTackle(level)
		assert.NoError(t, err)
		assert.Equal(t, voltTackle, pikachu.Move2)

		thunderstorm := NewThunderstorm(level)
		assert.Equal(t, thunderstorm, pikachu.UniteMove)

		// Assert buffs and additionalDamage are empty
		assert.Equal(t, stats.NewBuffs(), pikachu.Buffs)
		assert.Equal(t, attack.NewAllAdditionalDamage(), pikachu.AllAdditionalDamage)
	})
}

func TestPikachu_Attack(t *testing.T) {
	level := 9
	pikachu := setUpPikachu(t, level)
	enemy := &enemy.DefaultEnemy{Wild: false}

	t.Run("attack with basic attacks", func(t *testing.T) {
		// Basic attack 1
		damage, err := pikachu.Attack(attack.BasicAttackOption, enemy, 0)
		assert.NoError(t, err)
		assert.Equal(t, float64(213), damage.DamageDealt)
		assert.Equal(t, attack.BasicAttackOption, damage.AttackOption)
		assert.Equal(t, attack.PhysicalAttack, damage.AttackType)

		// Basic attack 2
		damage, err = pikachu.Attack(attack.BasicAttackOption, enemy, 1000)
		assert.NoError(t, err)
		assert.Equal(t, float64(213), damage.DamageDealt)
		assert.Equal(t, attack.BasicAttackOption, damage.AttackOption)
		assert.Equal(t, attack.PhysicalAttack, damage.AttackType)

		// Boosted attack
		damage, err = pikachu.Attack(attack.BasicAttackOption, enemy, 2000)
		assert.NoError(t, err)
		assert.Equal(t, 438.84, damage.DamageDealt)
		assert.Equal(t, attack.BasicAttackOption, damage.AttackOption)
		assert.Equal(t, attack.SpecialAttack, damage.AttackType)
	})
	t.Run("attack with move 1", func(t *testing.T) {
		// Attack with Electro Ball
		damage, err := pikachu.Attack(attack.Move1, enemy, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1005.88, damage.DamageDealt)
		assert.Equal(t, attack.Move1, damage.AttackOption)
		assert.Equal(t, attack.SpecialAttack, damage.AttackType)
	})
	t.Run("attack with move 2", func(t *testing.T) {
		// Attack with Volt tackle
		damage, err := pikachu.Attack(attack.Move2, enemy, 0)
		assert.NoError(t, err)
		assert.Equal(t, 222.52, damage.DamageDealt)
		assert.Equal(t, attack.Move2, damage.AttackOption)
		assert.Equal(t, attack.SpecialAttack, damage.AttackType)
	})
	t.Run("attack with unite move", func(t *testing.T) {
		damage, err := pikachu.Attack(attack.UniteMove, enemy, 0)
		assert.NoError(t, err)
		assert.Equal(t, attack.OverTimeDamage{
			Damage:          774.82,
			DamageFrequency: 750,
			DurationEnd:     3000,
		}, damage.OvertimeDamage)
		assert.Equal(t, attack.UniteMove, damage.AttackOption)
		assert.Equal(t, attack.SpecialAttack, damage.AttackType)
	})
}

func TestPikachu_ActivateBattleItem(t *testing.T) {
	level := 9
	pikachu := setUpPikachu(t, level)

	t.Run("Activate battle item", func(t *testing.T) {
		pikachu.ActivateBattleItem(0)
		assert.True(t, pikachu.Buffs[stats.BattleItemBuff] != stats.Buff{})
	})
}

func TestPikachu_GetAvailableActions(t *testing.T) {
	t.Run("All actions available", func(t *testing.T) {
		level := 9
		pikachu := setUpPikachu(t, level)

		// Assert actions are as expected
		availableActions, isBattleItemAvailable, err := pikachu.GetAvailableActions(0)
		assert.NoError(t, err)
		assert.Equal(t, []attack.Option{
			attack.BasicAttackOption,
			attack.Move1,
			attack.Move2,
			attack.UniteMove,
		}, availableActions)
		assert.True(t, isBattleItemAvailable)
	})
	t.Run("Only basic attack available", func(t *testing.T) {
		level := 9
		pikachu := setUpPikachu(t, level)

		// Use all attacks
		pikachu.ActivateBattleItem(0)
		pikachu.Attack(attack.Move1, &enemy.DefaultEnemy{}, 0)
		pikachu.Attack(attack.Move2, &enemy.DefaultEnemy{}, 1000)
		pikachu.Attack(attack.UniteMove, &enemy.DefaultEnemy{}, 2000)

		// Assert actions are as expected
		availableActions, isBattleItemAvailable, err := pikachu.GetAvailableActions(2000)
		assert.NoError(t, err)
		assert.Equal(t, []attack.Option{
			attack.BasicAttackOption,
		}, availableActions)
		assert.False(t, isBattleItemAvailable)
	})
}
