package wildpokemon

type RegielekiEnemy struct {
	WildPokemon
}

func NewRegielekiEnemy(remainingTime int) (regielekiEnemy *RegielekiEnemy, err error) {
	regielekiData, err := fetchWildPokemonData(RegielekiEnemyName, remainingTime)
	if err != nil {
		return nil, err
	}
	regielekiEnemy = &RegielekiEnemy{
		regielekiData,
	}
	return
}
