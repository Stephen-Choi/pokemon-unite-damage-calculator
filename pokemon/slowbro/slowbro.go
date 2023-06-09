package slowbro

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/datafetcher"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon/general_pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	SlowbroName = "slowbro"
)

type MoveName string

const (
	WaterGunName    MoveName = "water-gun"
	ScaldName       MoveName = "scald"
	SurfName        MoveName = "surf"
	SlackOffName    MoveName = "slack-off"
	AmnesiaName     MoveName = "amnesia"
	TelekinesisName MoveName = "telekinesis"
)

var move1Set = []MoveName{
	WaterGunName,
	ScaldName,
	SurfName,
}

var move2Set = []MoveName{
	SlackOffName,
	AmnesiaName,
	TelekinesisName,
}

type Slowbro struct {
	general_pokemon.GeneralPokemon
}

func NewSlowbro(level int, move1Name string, move2Name string, heldItems []helditems.HeldItem, battleItem battleitems.BattleItem, emblems *stats.Stats) (p *Slowbro, err error) {
	// Get pokemon stats
	pokemonStats, err := datafetcher.FetchPokemonStats(SlowbroName, level)
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

	p = &Slowbro{
		general_pokemon.GeneralPokemon{
			Name:                SlowbroName,
			Stats:               pokemonStats,
			BasicAttack:         NewBasicAttack(),
			Move1:               move1,
			Move2:               move2,
			UniteMove:           NewSlowbeam(level),
			Passive:             NewPassive(),
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
	case WaterGunName:
		move, err = NewWaterGun(level)
	case ScaldName:
		move, err = NewScald(level)
	case SurfName:
		move, err = NewSurf(level)
	case SlackOffName:
		move, err = NewSlackOff(level)
	case AmnesiaName:
		move, err = NewAmnesia(level)
	case TelekinesisName:
		move, err = NewTelekinesis(level)
	default:
		err = errors.New("invalid move name")
	}
	return
}
