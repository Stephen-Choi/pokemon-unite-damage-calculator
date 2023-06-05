package wildpokemon

import (
	"errors"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

const (
	RayquazaName         = "rayquaza"
	RegielekiNeutralName = "regieleki_neutral"
	RegielekiEnemyName   = "regieleki_enemy"
	BottomRegis          = "regis"
)

var ObjectivePokemonNames = []string{
	RayquazaName,
	RegielekiNeutralName,
	RegielekiEnemyName,
	BottomRegis,
}

var validRayquazaRemainingTimes = []int{
	120,
	100,
	70,
	40,
	10,
}

var validRegiRemainingTimes = []int{
	420,
	400,
	370,
	340,
	310,
	280,
	250,
	220,
	190,
	160,
	130,
}

// IsValidObjectivePokemon checks if a objective pokemon name is valid and can exists at a given remaining time
func IsValidObjectivePokemon(pokemonName string, remainingTime int) error {
	pokemonIsAnObjective := false
	for _, objectivePokemon := range ObjectivePokemonNames {
		if pokemonName == objectivePokemon {
			pokemonIsAnObjective = true
			break
		}
	}
	if !pokemonIsAnObjective {
		return errors.New("invalid objective pokemon name")
	}

	switch pokemonName {
	case RayquazaName:
		for _, validRemainingTime := range validRayquazaRemainingTimes {
			if remainingTime == validRemainingTime {
				return nil
			}
		}
	case RegielekiNeutralName, RegielekiEnemyName, BottomRegis:
		for _, validRemainingTime := range validRegiRemainingTimes {
			if remainingTime == validRemainingTime {
				return nil
			}
		}
	}

	return errors.New("invalid remaining time")
}

type WildPokemon struct {
	RemainingTime int     `json:"time_remaining"`
	StartingHp    float64 // Useful for certain effects
	stats.Stats
}

func (w *WildPokemon) IsWild() bool {
	return true
}

func (w *WildPokemon) GetStats() stats.Stats {
	return w.Stats
}

func (w *WildPokemon) GetRemainingHealth() float64 {
	return w.Stats.Hp
}

func (w *WildPokemon) GetMissingHealth() float64 {
	return w.StartingHp - w.Stats.Hp
}

func (w *WildPokemon) IsDefeated() bool {
	return w.GetRemainingHealth() <= 0
}

func (w *WildPokemon) ApplyDamage(damageTaken float64) {
	w.Stats.Hp = w.Stats.Hp - damageTaken
}

type JsonStats struct {
	RemainingTime int `json:"time_remaining"`
	Hp            int `json:"hp"`
	Attack        int `json:"attack"`
	Defense       int `json:"defense"`
	SpDefense     int `json:"sp_defense"`
}

func (j *JsonStats) ToStats() stats.Stats {
	return stats.Stats{
		Hp:             float64(j.Hp),
		Attack:         float64(j.Attack),
		Defense:        float64(j.Defense),
		SpecialDefense: float64(j.SpDefense),
	}
}
