package wildpokemon

type RegielekiNeutral struct {
	WildPokemon
}

func NewRegielekiNeutral(remainingTime int) (regielekiNeutral *RegielekiNeutral, err error) {
	regielekiData, err := fetchWildPokemonData(RegielekiNeutralName, remainingTime)
	if err != nil {
		return nil, err
	}
	regielekiNeutral = &RegielekiNeutral{
		regielekiData,
	}
	return
}
