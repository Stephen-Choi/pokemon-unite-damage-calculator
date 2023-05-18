package pokemon

import "github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"

const (
	absolName           = "absol"
	chandelureName      = "chandelure"
	decidueyeName       = "decidueye"
	espeonName          = "espeon"
	greninjaName        = "greninja"
	mrMimeName          = "mr_mime"
	talonflameName      = "talonflame"
	zacianName          = "zacian"
	aegislashName       = "aegislash"
	charizardName       = "charizard"
	delphoxName         = "delphox"
	garchompName        = "garchomp"
	hoopaName           = "hoopa"
	pikachuName         = "pikachu"
	trevenantName       = "trevenant"
	zeraoraName         = "zeraora"
	alolanNinetalesName = "alolan_ninetales"
	cinderaceName       = "cinderace"
	dodrioName          = "dodrio"
	gardevoirName       = "gardevoir"
	laprasName          = "lapras"
	sableyeName         = "sableye"
	tsareenaName        = "tsareena"
	zoroarkName         = "zoroark"
	azumarillName       = "azumarill"
	clefableName        = "clefable"
	dragapultName       = "dragapult"
	gengarName          = "gengar"
	lucarioName         = "lucario"
	scizorName          = "scizor"
	tyranitarName       = "tyranitar"
	blastoiseName       = "blastoise"
	comfeyName          = "comfey"
	dragoniteName       = "dragonite"
	glaceonName         = "glaceon"
	machampName         = "machamp"
	slowbroName         = "slowbro"
	urshifuName         = "urshifu"
	blisseyName         = "blissey"
	cramorantName       = "cramorant"
	duraludonName       = "duraludon"
	goodraName          = "goodra"
	mamoswineName       = "mamoswine"
	snorlaxName         = "snorlax"
	venusaurName        = "venusaur"
	buzzwoleName        = "buzzwole"
	crustleName         = "crustle"
	eldegossName        = "eldegoss"
	greedentName        = "greedent"
	mewName             = "mew"
	sylveonName         = "sylveon"
	wigglytuffName      = "wigglytuff"
)

var PlayablePokemons = []string{
	absolName,
	chandelureName,
	decidueyeName,
	espeonName,
	greninjaName,
	mrMimeName,
	talonflameName,
	zacianName,
	aegislashName,
	charizardName,
	delphoxName,
	garchompName,
	hoopaName,
	pikachuName,
	trevenantName,
	zeraoraName,
	alolanNinetalesName,
	cinderaceName,
	dodrioName,
	gardevoirName,
	laprasName,
	sableyeName,
	tsareenaName,
	zoroarkName,
	azumarillName,
	clefableName,
	dragapultName,
	gengarName,
	lucarioName,
	scizorName,
	tyranitarName,
	blastoiseName,
	comfeyName,
	dragoniteName,
	glaceonName,
	machampName,
	slowbroName,
	urshifuName,
	blisseyName,
	cramorantName,
	duraludonName,
	goodraName,
	mamoswineName,
	snorlaxName,
	venusaurName,
	buzzwoleName,
	crustleName,
	eldegossName,
	greedentName,
	mewName,
	sylveonName,
	wigglytuffName,
}

// IsPokemonPlayable checks if a pokemon is playable
func IsPokemonPlayable(pokemonName string) bool {
	for _, availablePokemon := range PlayablePokemons {
		if pokemonName == availablePokemon {
			return true
		}
	}
	return false
}

// Pokemon is an interface for all pokemon
type Pokemon interface {
	GetAvailableAttacks() (availableAttacks []attack.AttackOption, err error)                                          // Get the available attacks for a pokemon that are not on cooldown
	Attack(attack attack.AttackOption) (damageDealt int, additionalStatusEffects []attack.StatusConditions, err error) // Get the attack dealt by a pokemon's attack and possible status effects
}
