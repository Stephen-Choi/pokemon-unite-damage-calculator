package datafetcher

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Copied over from pokemon package to avoid cyclical dependency
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
}

// Test_fetchPokemonStats tests the FetchPokemonStats function
func Test_fetchPokemonStats(t *testing.T) {
	for _, playablePokemon := range PlayablePokemons {
		for level := 1; level < maxLevel; level++ {
			_, err := FetchPokemonStats(playablePokemon, level)
			assert.NoError(t, err, fmt.Sprintf("failed fetching pokemon: %s at level : %d", playablePokemon, level))
		}
	}
}
