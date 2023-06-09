package damagecalculator

import (
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/attack"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/enemy"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/stats"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	CannotAct     = "cannot act" // still in attack animation or attack lag
	UseBattleItem = "use battle item"
)

type PokemonActionResult struct {
	ActionName       string  `json:"action-name"` // Can be either an attack, the use of a battle item, or unable to act
	BaseDamage       float64 `json:"base-damage,omitempty"`
	CritDamage       float64 `json:"crit-damage,omitempty"`
	AdditionalDamage float64 `json:"additional-damage,omitempty"`
	OvertimeDamage   float64 `json:"overtime-damage,omitempty"`
	ExecutionDamage  float64 `json:"execution-damage-dealt,omitempty"`
	TotalDamageDealt float64 `json:"total-damage-dealt"`
}

type State struct {
	PokemonActions          map[string]PokemonActionResult        `json:"pokemon-actions"`
	PokemonBuffs            map[string]stats.Buffs                `json:"pokemon-buffs"`
	PokemonAdditionalDamage map[string]attack.AllAdditionalDamage `json:"pokemon-additional-damage"`
	OvertimeDamage          map[string][]attack.OverTimeDamage    `json:"overtime-damage"`
	PokemonTeamBuffs        stats.Buffs                           `json:"pokemon-team-buffs"`
	InflictedDebuffs        []attack.Debuff                       `json:"inflicted-debuffs"`
	EnemyHealth             float64                               `json:"enemy-health"`
}

// StateLog is a map of time to state details
type StateLog map[string]State

type Result struct {
	TotalTime float64  `json:"total_time"`
	StateLog  StateLog `json:"state_log"`
}

type DamageCalculator struct {
	attackingPokemon          map[string]pokemon.Pokemon
	attackingPokemonTeamBuff  stats.Buffs
	enemyPokemon              enemy.Pokemon
	inflictedDebuffs          []attack.Debuff
	overtimeDamageByPokemon   map[string][]attack.OverTimeDamage
	timeOfNextAvailableAction map[string]float64         // timeOfNextAvailableAction is a map of pokemon name to the time at which the next action is available
	pokemonPrevAction         map[string]attack.Option   // pokemonPrevAction is a map of pokemon name to its previous action
	attacksThatCanCrit        map[string][]attack.Option // attacksThatCanCrit is a map of pokemon name to its attacks that can crit

	elapsedTime float64
}

func NewDamageCalculator(attackingPokemon map[string]pokemon.Pokemon, enemyPokemon enemy.Pokemon, teamBuffs stats.Buffs, elapsedTime float64) *DamageCalculator {
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
		pokemonPrevAction:         make(map[string]attack.Option),
		inflictedDebuffs:          []attack.Debuff{},
		elapsedTime:               elapsedTime,
	}
}

// CalculateRip calculates how long it takes to defeat an enemy pokemon
// NOTE: assumes enemy will not attack and defeat any of the attacking pokemon
// TODO: set up an algorithm to determine the best course of action to achieve the most efficient rip
func (d *DamageCalculator) CalculateRip() (Result, error) {
	// Set up state log to capture all states during the rip calculation
	stateLog := make(StateLog)
	startingRipTime := d.elapsedTime

	// Keep calculating until the enemy is defeated
	for !d.enemyPokemon.IsDefeated() {
		// Elapse time (skip first iteration)
		if len(stateLog) != 0 {
			d.elapseTime()
		}

		// Set up new state for the state log
		state := State{
			PokemonActions:          make(map[string]PokemonActionResult),
			PokemonBuffs:            make(map[string]stats.Buffs),
			PokemonAdditionalDamage: make(map[string]attack.AllAdditionalDamage),
		}

		// Randomly select a pokemon to act (since it's a map, range will be random)
		for _, attackingPokemon := range d.attackingPokemon {
			attackingPokemon.ClearExpiredEffects(d.elapsedTime)
			attackingPokemonName := attackingPokemon.GetName()

			// Check what moves are available
			availableAttacks, isBattleItemAvailable, err := attackingPokemon.GetAvailableActions(d.elapsedTime)
			if err != nil {
				panic(err) // TODO: handle error
			}

			// If battle item is available, default to always activating it (take's up this pokemon's action)
			// Note: this can be activated during a move, so perform this before checking if the pokemon can act
			if isBattleItemAvailable {
				attackingPokemon.ActivateBattleItem(d.elapsedTime)
				state.PokemonActions[attackingPokemonName] = PokemonActionResult{ActionName: UseBattleItem}

				// TODO REFACTOR THIS (BEING USED IN 3 PLACES)

				// Need to make a copy of the buffs because the returned values are pointers
				pokemonBuffs := attackingPokemon.GetBuffs()
				pokemonBuffsCopy := make(map[string]stats.Buff)
				for pokemonName, buff := range pokemonBuffs {
					pokemonBuffsCopy[pokemonName] = buff
				}
				state.PokemonBuffs[attackingPokemonName] = pokemonBuffsCopy

				// Need to make a copy of the additional damage because the returned values are pointers
				pokemonAllAdditionalDamage := attackingPokemon.GetAllAdditionalDamage()

				allAdditionalDamage := make(map[string]attack.AdditionalDamage)
				for additionalDamageName, additionalDamage := range pokemonAllAdditionalDamage {
					allAdditionalDamage[additionalDamageName] = additionalDamage
				}
				state.PokemonAdditionalDamage[attackingPokemonName] = allAdditionalDamage
				continue
			}

			// Check if the pokemon can act
			if !d.canPokemonAct(attackingPokemonName, availableAttacks) {

				// TODO BUG FIX:
				// OVERTIME DAMAGE NEEDS TO OCCUR EVEN IF POKEMON IS IN ANIMATION
				state.PokemonActions[attackingPokemonName] = PokemonActionResult{ActionName: CannotAct}

				// TODO REFACTOR THIS (BEING USED IN 3 PLACES)

				// Need to make a copy of the buffs because the returned values are pointers
				pokemonBuffs := attackingPokemon.GetBuffs()
				pokemonBuffsCopy := make(map[string]stats.Buff)
				for pokemonName, buff := range pokemonBuffs {
					pokemonBuffsCopy[pokemonName] = buff
				}
				state.PokemonBuffs[attackingPokemonName] = pokemonBuffsCopy

				// Need to make a copy of the additional damage because the returned values are pointers
				pokemonAllAdditionalDamage := attackingPokemon.GetAllAdditionalDamage()

				allAdditionalDamage := make(map[string]attack.AdditionalDamage)
				for additionalDamageName, additionalDamage := range pokemonAllAdditionalDamage {
					allAdditionalDamage[additionalDamageName] = additionalDamage
				}
				state.PokemonAdditionalDamage[attackingPokemonName] = allAdditionalDamage
				continue
			}

			// Perform the attack
			bestAction := determineBestAction(availableAttacks)
			attackResult, err := attackingPokemon.Attack(bestAction, d.enemyPokemon, d.elapsedTime)
			if err != nil {
				panic(err) // TODO: handle error
			}

			// Check if any overtime damage occurred
			if attackResult.OvertimeDamage.Exists() {
				d.overtimeDamageByPokemon[attackingPokemonName] = append(d.overtimeDamageByPokemon[attackingPokemonName], attackResult.OvertimeDamage)
			}

			// TODO: if any team buffs occured, add them here...

			// Add debuffs to inflictedDebuffs
			if len(attackResult.Debuffs) > 0 {
				d.inflictedDebuffs = append(d.inflictedDebuffs, attackResult.Debuffs...)
			}

			// Check for crit
			var critDamage float64
			if d.shouldCrit(attackingPokemon.GetName(), bestAction, d.elapsedTime) {
				critDamageMultiplier := attackingPokemon.GetStats().CriticalHitDamage
				critDamage = (attackResult.BaseDamageDealt * critDamageMultiplier) - attackResult.BaseDamageDealt
			}

			// Determine damage to be taken by enemy
			enemyStatsAfterDebuffs := attack.ApplyDebuffs(attackingPokemonName, d.enemyPokemon, d.inflictedDebuffs)
			baseDamageDealt := calculateDamageDealt(attackResult.BaseDamageDealt, enemyStatsAfterDebuffs, attackResult.AttackType)
			critDamageDealt := calculateDamageDealt(critDamage, enemyStatsAfterDebuffs, attackResult.AttackType)
			additionalDamageDealt := calculateDamageDealt(attackResult.AdditionalDamageDealt, enemyStatsAfterDebuffs, attackResult.AttackType)
			overtimeDamageDealt := d.calculateOvertimeDamage(attackingPokemonName, enemyStatsAfterDebuffs, d.elapsedTime)
			totalDamageDealt := baseDamageDealt + critDamageDealt + additionalDamageDealt + overtimeDamageDealt

			// Apply damage to enemy's health
			d.enemyPokemon.ApplyDamage(totalDamageDealt)

			// Apply any execution damage if required (applies after damage is dealt from the move)
			// Ref on how to calculate execution damage: https://unite-db.com/faq/elementary-mechanics#Missing-HP
			var executionDamageDealt float64
			if attackResult.ExecutionPercentDamage.Exists() {
				var executionDamage float64
				if attackResult.ExecutionPercentDamage.CappedDamage != 0 {
					executionDamage = math.Min(d.enemyPokemon.GetMissingHealth()*attackResult.ExecutionPercentDamage.Percent, attackResult.ExecutionPercentDamage.CappedDamage)
				} else {
					executionDamage = d.enemyPokemon.GetMissingHealth() * attackResult.ExecutionPercentDamage.Percent
				}

				executionDamageDealt = calculateDamageDealt(executionDamage, enemyStatsAfterDebuffs, attackResult.AttackType)
				d.enemyPokemon.ApplyDamage(executionDamageDealt)
				totalDamageDealt += executionDamageDealt
			}

			// Update state log for this pokemon
			state.PokemonActions[attackingPokemonName] = PokemonActionResult{
				ActionName:       attackResult.AttackName,
				BaseDamage:       baseDamageDealt,
				CritDamage:       critDamageDealt,
				AdditionalDamage: additionalDamageDealt,
				OvertimeDamage:   overtimeDamageDealt,
				ExecutionDamage:  executionDamageDealt,
				TotalDamageDealt: totalDamageDealt,
			}

			// Need to make a copy of the buffs because the returned values are pointers
			pokemonBuffs := attackingPokemon.GetBuffs()
			pokemonBuffsCopy := make(map[string]stats.Buff)
			for pokemonName, buff := range pokemonBuffs {
				pokemonBuffsCopy[pokemonName] = buff
			}
			state.PokemonBuffs[attackingPokemonName] = pokemonBuffsCopy

			// Need to make a copy of the additional damage because the returned values are pointers
			pokemonAllAdditionalDamage := attackingPokemon.GetAllAdditionalDamage()

			allAdditionalDamage := make(map[string]attack.AdditionalDamage)
			for additionalDamageName, additionalDamage := range pokemonAllAdditionalDamage {
				allAdditionalDamage[additionalDamageName] = additionalDamage
			}
			state.PokemonAdditionalDamage[attackingPokemonName] = allAdditionalDamage

			// Set delay for next action for the attacking pokemon
			d.setActionDelay(attackingPokemonName, attackResult, d.elapsedTime)
			d.pokemonPrevAction[attackingPokemonName] = attackResult.AttackOption
		}

		// Set copy of overtime damage at this elapsed time
		pokemonOvertimeDamage := make(map[string][]attack.OverTimeDamage)
		for pokemonName, overtimeDamage := range d.overtimeDamageByPokemon {
			pokemonOvertimeDamage[pokemonName] = overtimeDamage
		}
		state.OvertimeDamage = pokemonOvertimeDamage

		// Add entry to StateLog
		state.PokemonTeamBuffs = d.attackingPokemonTeamBuff
		state.InflictedDebuffs = d.inflictedDebuffs
		state.EnemyHealth = d.enemyPokemon.GetRemainingHealth()
		stateLog[strconv.FormatFloat(d.elapsedTime, 'f', -1, 64)] = state
	}

	// Return result
	return Result{
		StateLog:  stateLog,
		TotalTime: d.elapsedTime - startingRipTime,
	}, nil

	return Result{}, nil
}

// setActionDelay sets the delay for the next action for the attacking pokemon
// Basic attacks have buckets that determine the next actionable frame
// There is currently no frame data for skill moves to determine their exact duration, so arbitrarily choosing a 750 millisecond delay
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
		attackingPokemonAttackSpeed := d.attackingPokemon[attackingPokemonName].GetStats().AttackSpeed
		actionDelay = attack.GetDelayForAttackSpeed(attackingPokemonAttackSpeed)
	} else {
		// Set delay to 750 second for any skill/unite move
		actionDelay = 750
	}
	d.timeOfNextAvailableAction[attackingPokemonName] = elapsedTime + actionDelay
}

func (d *DamageCalculator) canPokemonAct(pokemonName string, availableAttacks []attack.Option) bool {
	// If prev move was basic attack, and only available move is basic attack, pokemon must wait through the attack speed delay
	prevMoveWasBasicAttack := d.pokemonPrevAction[pokemonName] == attack.BasicAttackOption

	prevAndCurrentAttacksAreBasic := prevMoveWasBasicAttack && len(availableAttacks) == 1 && availableAttacks[0] == attack.BasicAttackOption
	if prevAndCurrentAttacksAreBasic {
		return d.timeOfNextAvailableAction[pokemonName] <= d.elapsedTime
	}

	// If prev move was a basic attack and a skill move is available, pokemon can act immediately
	if prevMoveWasBasicAttack && len(availableAttacks) > 1 {
		return true
	}

	// Prev move was not a basic attack, pokemon must wait through the skill animation delay
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
func (d *DamageCalculator) calculateOvertimeDamage(pokemonName string, enemyStats stats.Stats, elapsedTime float64) float64 {
	totalOvertimeDamageDealt := 0.0
	allOvertimeDamage := d.overtimeDamageByPokemon[pokemonName]
	for index, overtimeDamage := range allOvertimeDamage {
		if shouldApplyOvertimeDamage(overtimeDamage, elapsedTime) {
			totalOvertimeDamageDealt += calculateDamageDealt(overtimeDamage.BaseDamage, enemyStats, overtimeDamage.AttackType)
			d.overtimeDamageByPokemon[pokemonName][index].LastInflictedDamageTime = elapsedTime
		}
	}

	return totalOvertimeDamageDealt
}

// shouldApplyOvertimeDamage checks if the overtime damage should be applied given its damage frequency
func shouldApplyOvertimeDamage(overtimeDamage attack.OverTimeDamage, elapsedTime float64) bool {
	// Return true if no damage has been inflicted yet
	if overtimeDamage.LastInflictedDamageTime == 0 {
		return true
	}

	return overtimeDamage.LastInflictedDamageTime+overtimeDamage.DamageFrequency <= elapsedTime
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
	critChance := d.attackingPokemon[pokemonName].GetStats().CriticalHitChance
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

// calculateDamageDealt calculates the damage taken by the enemy pokemon
func calculateDamageDealt(attackDamage float64, enemyStats stats.Stats, attackType attack.Type) float64 {
	var damageDealt float64
	if attackType == attack.PhysicalAttack {
		damageDealt = attackDamage * 600 / (600 + enemyStats.Defense)
	} else {
		damageDealt = attackDamage * 600 / (600 + enemyStats.SpecialDefense)
	}
	return math.Floor(damageDealt)
}
