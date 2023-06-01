package pikachu

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/datafetcher"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon/general_pokemon"
	stats "github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	PikachuName = "pikachu"
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
	general_pokemon.GeneralPokemon
}

func NewPikachu(level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem, emblems *stats.Stats) (p *Pikachu, err error) {
	// Get pokemon stats
	pokemonStats, err := datafetcher.FetchPokemonStats(PikachuName, level)
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
		general_pokemon.GeneralPokemon{
			Name:                PikachuName,
			Stats:               pokemonStats,
			BasicAttack:         NewBasicAttack(),
			Move1:               move1,
			Move2:               move2,
			UniteMove:           NewThunderstorm(level),
			HeldItems:           heldItems,
			BattleItem:          battleItem,
			Buffs:               stats.NewBuffs(),
			AllAdditionalDamage: attack.NewAllAdditionalDamage(),
		},
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
