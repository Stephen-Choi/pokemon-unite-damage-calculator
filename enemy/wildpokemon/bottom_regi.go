package wildpokemon

type BottomRegi struct {
	WildPokemon
}

func NewBottomRegi(remainingTime int) (bottomRegi *BottomRegi, err error) {
	bottomRegiData, err := fetchWildPokemonData(BottomRegis, remainingTime)
	if err != nil {
		return nil, err
	}
	bottomRegi = &BottomRegi{
		bottomRegiData,
	}
	return
}
