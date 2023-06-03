package damage_calculator

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math"
	"math/rand"
	"time"
)

type State struct {
	PokemonActions   map[string]attack.Option           `json:"pokemon_actions"`
	PokemonBuffs     map[string]stats.Buffs             `json:"pokemon_buffs"`
	PokemonTeamBuffs stats.Buffs                        `json:"pokemon_team_buffs"`
	InflictedDebuffs []attack.Debuff                    `json:"inflicted_debuffs"`
	OvertimeDamage   map[string][]attack.OverTimeDamage `json:"overtime_damage"`
	EnemyHealth      float64                            `json:"enemy_health"`
}

// StateLog is a map of time to state details
type StateLog map[float64]State

type Result struct {
	TotalTime float64 `json:"total_time"`
	StateLog  `json:"state_log"`
}

type DamageCalculator struct {
	attackingPokemon          map[string]pokemon.Pokemon
	attackingPokemonTeamBuff  stats.Buffs
	enemyPokemon              enemy.Pokemon
	inflictedDebuffs          []attack.Debuff
	overtimeDamageByPokemon   map[string][]attack.OverTimeDamage
	timeOfNextAvailableAction map[string]float64         // timeOfNextAvailableAction is a map of pokemon name to the time at which the next action is available
	attacksThatCanCrit        map[string][]attack.Option // attacksThatCanCrit is a map of pokemon name to its attacks that can crit

	elapsedTime float64
}

func NewDamageCalculator(attackingPokemon map[string]pokemon.Pokemon, enemyPokemon enemy.Pokemon, teamBuffs stats.Buffs) *DamageCalculator {
	// Set up attacks that can crit
	attacksThatCanCrit := make(map[string][]attack.Option)
	for _, pokemon := range attackingPokemon {
		attacksThatCanCrit[pokemon.GetName()] = pokemon.GetMovesThatCanCrit()
	}

	return &DamageCalculator{
		attackingPokemon:          attackingPokemon,
		attackingPokemonTeamBuff:  teamBuffs,
		enemyPokemon:              enemyPokemon,
		attacksThatCanCrit:        attacksThatCanCrit,
		overtimeDamageByPokemon:   make(map[string][]attack.OverTimeDamage),
		timeOfNextAvailableAction: make(map[string]float64),
	}
}

// CalculateRip calculates how long it takes to defeat an enemy pokemon
// NOTE: assumes enemy will not attack and defeat any of the attacking pokemon
// TODO: set up an algorithm to determine the best course of action to achieve the most efficient rip
func (d *DamageCalculator) CalculateRip() (Result, error) {
	// Set up state log to capture all states during the rip calculation
	stateLog := make(StateLog)

	// Keep calculating until the enemy is defeated
	for !d.enemyPokemon.IsDefeated() {
		// Elapse time (skip first iteration)
		if len(stateLog) != 0 {
			d.elapseTime()
		}

		// Set up new state for the state log
		state := State{
			PokemonActions: make(map[string]attack.Option),
			PokemonBuffs:   make(map[string]stats.Buffs),
		}

		// Randomly select a pokemon to act (since it's a map, range will be random)
		for _, attackingPokemon := range d.attackingPokemon {
			attackingPokemonName := attackingPokemon.GetName()

			// Check if the pokemon can act
			if !d.canPokemonAct(attackingPokemonName) {
				state.PokemonActions[attackingPokemonName] = attack.CannotAct
				continue
			}

			// If pokemon can act, determine the best action to take
			availableAttacks, isBattleItemAvailable, err := attackingPokemon.GetAvailableActions(d.elapsedTime)
			if err != nil {
				panic(err) // TODO: handle error
			}
			// If battle item is available, default to always activating it
			// TODO: check if battle item can be activated even during a move...
			if isBattleItemAvailable {
				attackingPokemon.ActivateBattleItem(d.elapsedTime)
			}
			// Perform the attack
			bestAction := determineBestAction(availableAttacks)
			attackResult, err := attackingPokemon.Attack(bestAction, d.enemyPokemon, d.elapsedTime)
			if err != nil {
				panic(err) // TODO: handle error
			}

			// Check if any overtime damage occured
			if attackResult.OvertimeDamage.Exists() {
				d.overtimeDamageByPokemon[attackingPokemonName] = append(d.overtimeDamageByPokemon[attackingPokemonName], attackResult.OvertimeDamage)
			}
			// Calculate total overtime damage to be inflicted
			overtimeDamageForPokemon := d.calculateOvertimeDamage(attackingPokemonName, d.elapsedTime)

			// TODO: if any team buffs occured, add them here...

			// Add debuffs to inflictedDebuffs
			if len(attackResult.Debuffs) > 0 {
				d.inflictedDebuffs = append(d.inflictedDebuffs, attackResult.Debuffs...)
			}

			// Check for crit
			if d.shouldCrit(attackingPokemon.GetName(), bestAction, d.elapsedTime) {
				critDamage := attackingPokemon.GetStats(d.elapsedTime).CriticalHitDamage
				attackResult.DamageDealt = attackResult.DamageDealt * critDamage // Applies crit damage
			}

			// Determine damage to be taken by enemy
			totalDamageDealt := attackResult.DamageDealt + overtimeDamageForPokemon
			enemyStatsAfterDebuffs := attack.ApplyDebuffs(attackingPokemonName, d.enemyPokemon, d.inflictedDebuffs)
			damageTakenByEnemy := calculateDamageTaken(totalDamageDealt, enemyStatsAfterDebuffs, attackResult.AttackType)

			// Apply damage to enemy's health
			d.enemyPokemon.ApplyDamage(damageTakenByEnemy)

			// Apply any execution damage if required (applies after damage is dealt from the move)
			// Ref on how to calculate execution damage: https://unite-db.com/faq/elementary-mechanics#Missing-HP
			if attackResult.ExecutionPercentDamage.Exists() {
				var executionDamage float64
				if attackResult.ExecutionPercentDamage.CappedDamage != 0 {
					executionDamage = math.Min(d.enemyPokemon.GetRemainingHealth()*attackResult.ExecutionPercentDamage.Percent, attackResult.ExecutionPercentDamage.CappedDamage)
				} else {
					executionDamage = d.enemyPokemon.GetRemainingHealth() * attackResult.ExecutionPercentDamage.Percent
				}

				executionDamageTakenByEnemy := calculateDamageTaken(executionDamage, enemyStatsAfterDebuffs, attackResult.AttackType)
				d.enemyPokemon.ApplyDamage(executionDamageTakenByEnemy)
			}

			// Update state log for this pokemon
			state.PokemonActions[attackingPokemonName] = bestAction
			state.PokemonBuffs[attackingPokemonName] = attackingPokemon.GetBuffs(d.elapsedTime)

			// Set delay for next action for the attacking pokemon
			d.setActionDelay(attackingPokemonName, attackResult, d.elapsedTime)
		}

		// Add entry to StateLog
		state.PokemonTeamBuffs = d.attackingPokemonTeamBuff
		state.InflictedDebuffs = d.inflictedDebuffs
		state.OvertimeDamage = d.overtimeDamageByPokemon
		state.EnemyHealth = d.enemyPokemon.GetRemainingHealth()
		stateLog[d.elapsedTime] = state
	}

	// Return result
	return Result{
		StateLog:  stateLog,
		TotalTime: d.elapsedTime,
	}, nil
}

// setActionDelay sets the delay for the next action for the attacking pokemon
// Basic attacks have buckets that determine the next actionable frame
// There is currently no frame data for skill moves to determine their exact duration, so arbitrarily choosing a 1 second delay
// TODO: if someone has frame data for skill moves, please update attack duration field for those moves
func (d *DamageCalculator) setActionDelay(attackingPokemonName string, attackResult attack.Result, elapsedTime float64) {
	attackDuration := attackResult.AttackDuration
	attackOption := attackResult.AttackOption

	// If attack duration is available, simply add that as the delay until the next action
	if attackDuration != 0 {
		d.timeOfNextAvailableAction[attackingPokemonName] = elapsedTime + attackDuration
		return
	}

	// If attack duration is not available, determine the delay based on the attack option
	var actionDelay float64
	if attackOption == attack.BasicAttackOption {
		attackingPokemonAttackSpeed := d.attackingPokemon[attackingPokemonName].GetStats(elapsedTime).AttackSpeed
		actionDelay = attack.GetDelayForAttackSpeed(attackingPokemonAttackSpeed)
	} else {
		// Set delay to 1 second for any skill/unite move
		actionDelay = 1000
	}
	d.timeOfNextAvailableAction[attackingPokemonName] = elapsedTime + actionDelay
}

func (d *DamageCalculator) canPokemonAct(pokemonName string) bool {
	return d.timeOfNextAvailableAction[pokemonName] <= d.elapsedTime
}

// determineBestAction determines the best action to take given the available attacks
// NOTE: starting this off very simple, will always prioritize a skill move if available
// The order of priority is: unite move > move 2 > move 1 > basic attack
func determineBestAction(availableAttacks []attack.Option) attack.Option {
	bestAction := attack.BasicAttackOption // Default to basic attack if no other options are available
	for _, attackOption := range availableAttacks {
		if attackOption == attack.UniteMove {
			return attackOption
		} else if attackOption == attack.Move2 {
			bestAction = attackOption
		} else if attackOption == attack.Move1 && bestAction != attack.Move2 {
			bestAction = attackOption
		}
	}
	return bestAction
}

// calculateOvertimeDamage calculates the damage dealt by overtime attacks
func (d *DamageCalculator) calculateOvertimeDamage(pokemonName string, elapsedTime float64) float64 {
	totalOvertimeDamage := 0.0
	allOvertimeDamage := d.overtimeDamageByPokemon[pokemonName]
	for _, overtimeDamage := range allOvertimeDamage {
		if shouldApplyOvertimeDamage(overtimeDamage, elapsedTime) {
			totalOvertimeDamage += overtimeDamage.Damage
		}
	}

	return totalOvertimeDamage
}

// shouldApplyOvertimeDamage checks if the overtime damage should be applied given its damage frequency
func shouldApplyOvertimeDamage(overtimeDamage attack.OverTimeDamage, elapsedTime float64) bool {
	return math.Mod(elapsedTime-overtimeDamage.DurationStart, overtimeDamage.DamageFrequency) == 0
}

// removeExpiredOvertimeDamage removes overtime damage that have expired
func (d *DamageCalculator) removeExpiredOvertimeDamage(elapsedTime float64) {
	for pokemonName, overtimeDamage := range d.overtimeDamageByPokemon {
		nonExpiredOvertimeDamage := []attack.OverTimeDamage{}
		for _, overtimeDamage := range overtimeDamage {
			if overtimeDamage.DurationEnd >= elapsedTime {
				nonExpiredOvertimeDamage = append(nonExpiredOvertimeDamage, overtimeDamage)
			}
		}
		d.overtimeDamageByPokemon[pokemonName] = nonExpiredOvertimeDamage
	}
}

// shouldCrit checks if an attack should crit
func (d *DamageCalculator) shouldCrit(pokemonName string, attackOption attack.Option, elapsedTime float64) bool {
	// Check if the attack can crit
	canCrit := false
	for _, attackThatCanCrit := range d.attacksThatCanCrit[pokemonName] {
		if attackThatCanCrit == attackOption {
			canCrit = true
		}
	}

	// If attack cannot crit, return early
	if !canCrit {
		return false
	}

	// If attack can crit, check if it crits
	critChance := d.attackingPokemon[pokemonName].GetStats(elapsedTime).CriticalHitChance
	return rollCrit(critChance)
}

// rollCrit rolls a crit based on the crit chance
func rollCrit(critChance float64) bool {
	rand.Seed(time.Now().UnixNano())

	// Scale the probability to an integer range (0-10000)
	scaledProbability := int(critChance * 100)

	// Generate a random number between 0 and 10000
	randomNumber := rand.Intn(10001)

	// Check if the random number falls within the desired probability range
	if randomNumber <= scaledProbability {
		return true
	}
	return false
}

// elapsedTime elapses time by 1/15 of a second which is the smallest unit of time in the game
func (d *DamageCalculator) elapseTime() {
	d.elapsedTime += 66.67

	// Remove overtime damage that have expired
	d.removeExpiredOvertimeDamage(d.elapsedTime)
}

// calculateDamageTaken calculates the damage taken by the enemy pokemon
func calculateDamageTaken(attackDamage float64, enemyStats stats.Stats, attackType attack.Type) float64 {
	var damageTaken float64
	if attackType == attack.PhysicalAttack {
		damageTaken = attackDamage * 600 / (600 + enemyStats.Defense)
	} else {
		damageTaken = attackDamage * 600 / (600 + enemyStats.SpecialDefense)
	}
	return damageTaken
}
