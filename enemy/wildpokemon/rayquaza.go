package wildpokemon

import "fmt"

type Rayquaza struct {
	WildPokemon
}

func NewRayquaza(remainingTime int) (rayquaza *Rayquaza, err error) {
	rayquazaData, err := fetchWildPokemonData(RayquazaName, remainingTime)
	if err != nil {
		return nil, err
	}

	fmt.Printf("wild pokemon: %+v\n", rayquazaData)

	rayquaza = &Rayquaza{
		rayquazaData,
	}
	return
}
