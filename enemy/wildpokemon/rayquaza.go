package wildpokemon

type Rayquaza struct {
	WildPokemon
}

func NewRayquaza(remainingTime int) (rayquaza *Rayquaza, err error) {
	rayquazaData, err := fetchWildPokemonData(RayquazaName, remainingTime)
	if err != nil {
		return nil, err
	}
	rayquaza = &Rayquaza{
		rayquazaData,
	}
	return
}
