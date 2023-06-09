package damagecalculator

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
)

// State represents the state of the rip calculation at a specific point in time
type State struct {
	PokemonActions          map[string]PokemonActionResult        `json:"pokemon-actions"`
	PokemonBuffs            map[string]stats.Buffs                `json:"pokemon-buffs"`
	PokemonAdditionalDamage map[string]attack.AllAdditionalDamage `json:"pokemon-additional-damage"`
	OvertimeDamage          map[string][]attack.OverTimeDamage    `json:"overtime-damage"`
	PokemonTeamBuffs        stats.Buffs                           `json:"pokemon-team-buffs"`
	InflictedDebuffs        []attack.Debuff                       `json:"inflicted-debuffs"`
	EnemyHealth             float64                               `json:"enemy-health"`
}

// setAttackingPokemonBuffs deep copies a pokemon's buffs into the state
func (s *State) setAttackingPokemonBuffs(attackingPokemon pokemon.Pokemon) {
	pokemonBuffs := attackingPokemon.GetBuffs()
	pokemonBuffsCopy := make(map[string]stats.Buff)
	for pokemonName, buff := range pokemonBuffs {
		pokemonBuffsCopy[pokemonName] = buff
	}
	s.PokemonBuffs[attackingPokemon.GetName()] = pokemonBuffsCopy
}

// setAllAdditionalDamage deep copies a pokemon's additional damage into the state
func (s *State) setAllAdditionalDamage(attackingPokemon pokemon.Pokemon) {
	pokemonAdditionalDamage := attackingPokemon.GetAllAdditionalDamage()
	pokemonAdditionalDamageCopy := make(map[string]attack.AdditionalDamage)
	for pokemonName, additionalDamage := range pokemonAdditionalDamage {
		pokemonAdditionalDamageCopy[pokemonName] = additionalDamage
	}
	s.PokemonAdditionalDamage[attackingPokemon.GetName()] = pokemonAdditionalDamageCopy
}

// setOvertimeDamage deep copies a pokemon's overtime damage into the state
func (s *State) setOvertimeDamage(overtimeDamageByPokemon map[string][]attack.OverTimeDamage) {
	pokemonOvertimeDamage := make(map[string][]attack.OverTimeDamage)
	for pokemonName, overtimeDamage := range overtimeDamageByPokemon {
		pokemonOvertimeDamage[pokemonName] = overtimeDamage
	}
	s.OvertimeDamage = pokemonOvertimeDamage
}

// setPokemonTeamBuffs deep copies a team's buffs into the state
func (s *State) setPokemonTeamBuffs(teamBuffs stats.Buffs) {
	teamBuffsCopy := stats.NewBuffs()
	for pokemonName, buff := range teamBuffs {
		teamBuffsCopy[pokemonName] = buff
	}
	s.PokemonTeamBuffs = teamBuffsCopy
}

// setInflictedDebuffs deep copies inflicted debuffs into the state
func (s *State) setInflictedDebuffs(inflictedDebuffs []attack.Debuff) {
	inflictedDebuffsCopy := []attack.Debuff{}
	for _, debuff := range inflictedDebuffs {
		inflictedDebuffsCopy = append(inflictedDebuffsCopy, debuff)
	}
	s.InflictedDebuffs = inflictedDebuffsCopy
}
