package pokemon

import "testing"

func Test_GetPokemon(t *testing.T) {
	t.Run("success - pikachu", func(t *testing.T) {
		pokemon, err := GetPokemon(PikachuName, 1, "thunder_shock", "electroweb", nil, nil)
		if err != nil {
			t.Fatalf("error getting pikachu: %v", err)
		}

		if pokemon.GetName() != PikachuName {
			t.Fatalf("expected pokemon name to be %s, got %s", PikachuName, pokemon.GetName())
		}
	})
	t.Run("error - invalid pokemon", func(t *testing.T) {
		_, err := GetPokemon("invalid pokemon", 1, "", "", nil, nil)
		if err == nil {
			t.Fatal("expected error getting invalid pokemon, got nil")
		}
	})
	t.Run("error - invalid moveset", func(t *testing.T) {
		_, err := GetPokemon("PikachuName", 1, "thunder_shock", "thunder_shock", nil, nil)
		if err == nil {
			t.Fatal("expected error getting invalid pokemon, got nil")
		}
	})
}
