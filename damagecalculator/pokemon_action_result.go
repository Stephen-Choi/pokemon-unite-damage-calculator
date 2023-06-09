package damagecalculator

const (
	CannotAct     = "cannot act" // still in attack animation or attack lag
	UseBattleItem = "use battle item"
)

// PokemonActionResult is the result of a pokemon's action
type PokemonActionResult struct {
	ActionName       string  `json:"action-name"` // Can be either an attack, the use of a battle item, or unable to act
	BaseDamage       float64 `json:"base-damage,omitempty"`
	CritDamage       float64 `json:"crit-damage,omitempty"`
	AdditionalDamage float64 `json:"additional-damage,omitempty"`
	OvertimeDamage   float64 `json:"overtime-damage,omitempty"`
	ExecutionDamage  float64 `json:"execution-damage-dealt,omitempty"`
	TotalDamageDealt float64 `json:"total-damage-dealt"`
}

func (p *PokemonActionResult) setTotalDamageDealt() float64 {
	p.TotalDamageDealt = p.BaseDamage + p.CritDamage + p.AdditionalDamage + p.OvertimeDamage + p.ExecutionDamage
	return p.TotalDamageDealt
}

// copyAttackResult populates a PokemonActionResult with the data returned from an attack
func (p *PokemonActionResult) copyAttackResult(attackActionResult PokemonActionResult) {
	p.ActionName = attackActionResult.ActionName
	p.BaseDamage = attackActionResult.BaseDamage
	p.CritDamage = attackActionResult.CritDamage
	p.AdditionalDamage = attackActionResult.AdditionalDamage
	p.ExecutionDamage = attackActionResult.ExecutionDamage
}
