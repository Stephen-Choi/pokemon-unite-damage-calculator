package pokemon

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const maxLevel = 15

const (
	AbsolName           = "absol"
	ChandelureName      = "chandelure"
	DecidueyeName       = "decidueye"
	EspeonName          = "espeon"
	GreninjaName        = "greninja"
	MrMimeName          = "mr_mime"
	TalonflameName      = "talonflame"
	ZacianName          = "zacian"
	AegislashName       = "aegislash"
	CharizardName       = "charizard"
	DelphoxName         = "delphox"
	GarchompName        = "garchomp"
	HoopaName           = "hoopa"
	PikachuName         = "pikachu"
	TrevenantName       = "trevenant"
	ZeraoraName         = "zeraora"
	AlolanNinetalesName = "alolan_ninetales"
	CinderaceName       = "cinderace"
	DodrioName          = "dodrio"
	GardevoirName       = "gardevoir"
	LaprasName          = "lapras"
	SableyeName         = "sableye"
	TsareenaName        = "tsareena"
	ZoroarkName         = "zoroark"
	AzumarillName       = "azumarill"
	ClefableName        = "clefable"
	DragapultName       = "dragapult"
	GengarName          = "gengar"
	LucarioName         = "lucario"
	ScizorName          = "scizor"
	TyranitarName       = "tyranitar"
	BlastoiseName       = "blastoise"
	ComfeyName          = "comfey"
	DragoniteName       = "dragonite"
	GlaceonName         = "glaceon"
	MachampName         = "machamp"
	SlowbroName         = "slowbro"
	UrshifuName         = "urshifu"
	BlisseyName         = "blissey"
	CramorantName       = "cramorant"
	DuraludonName       = "duraludon"
	GoodraName          = "goodra"
	MamoswineName       = "mamoswine"
	SnorlaxName         = "snorlax"
	VenusaurName        = "venusaur"
	BuzzwoleName        = "buzzwole"
	CrustleName         = "crustle"
	EldegossName        = "eldegoss"
	GreedentName        = "greedent"
	MewName             = "mew"
	SylveonName         = "sylveon"
	WigglytuffName      = "wigglytuff"
	UmbreonName         = "umbreon"
)

var PlayablePokemons = []string{
	AbsolName,
	ChandelureName,
	DecidueyeName,
	EspeonName,
	GreninjaName,
	MrMimeName,
	TalonflameName,
	ZacianName,
	AegislashName,
	CharizardName,
	DelphoxName,
	GarchompName,
	HoopaName,
	PikachuName,
	TrevenantName,
	ZeraoraName,
	AlolanNinetalesName,
	CinderaceName,
	DodrioName,
	GardevoirName,
	LaprasName,
	SableyeName,
	TsareenaName,
	ZoroarkName,
	AzumarillName,
	ClefableName,
	DragapultName,
	GengarName,
	LucarioName,
	ScizorName,
	TyranitarName,
	BlastoiseName,
	ComfeyName,
	DragoniteName,
	GlaceonName,
	MachampName,
	SlowbroName,
	UrshifuName,
	BlisseyName,
	CramorantName,
	DuraludonName,
	GoodraName,
	MamoswineName,
	SnorlaxName,
	VenusaurName,
	BuzzwoleName,
	CrustleName,
	EldegossName,
	GreedentName,
	MewName,
	SylveonName,
	WigglytuffName,
	UmbreonName,
}

var Defenders = []string{
	BlastoiseName,
	CrustleName,
	GoodraName,
	GreedentName,
	LaprasName,
	MamoswineName,
	SlowbroName,
	SnorlaxName,
	TrevenantName,
	UmbreonName,
}

var Supporters = []string{
	BlisseyName,
	ClefableName,
	ComfeyName,
	EldegossName,
	HoopaName,
	MrMimeName,
	SableyeName,
	WigglytuffName,
}

// isPokemonPlayable checks if a pokemon is playable
func isPokemonPlayable(pokemonName string) bool {
	for _, availablePokemon := range PlayablePokemons {
		if pokemonName == availablePokemon {
			return true
		}
	}
	return false
}

// isPokemonIsDefenderOrSupporter checks if a pokemon is a defender or supporter
func isPokemonIsDefenderOrSupporter(pokemonName string) bool {
	for _, defender := range Defenders {
		if pokemonName == defender {
			return true
		}
	}
	for _, supporter := range Supporters {
		if pokemonName == supporter {
			return true
		}
	}
	return false
}

// Pokemon is an interface for all playable pokemon
type Pokemon interface {
	GetName() string
	GetAvailableActions(elapsedTime float64) (availableAttacks []attack.Option, isBattleItemAvailable bool, err error) // Get the available actions for a pokemon
	Attack(attack attack.Option, enemyPokemon enemy.Pokemon, elapsedTime float64) (result attack.Result, err error)    // Get the attack dealt by a pokemon's attack and possible status effects
	GetMovesThatCanCrit() []attack.Option                                                                              // Get the list of moves that can crit
	ActivateBattleItem(elapsedTime float64)
	GetStats() stats.Stats
	GetBuffs() stats.Buffs
	GetAllAdditionalDamage() attack.AllAdditionalDamage
	ClearExpiredEffects(elapsedTime float64)
	AddAdditionalDamage(additionalDamageName string, additionalDamage attack.AdditionalDamage)
}
