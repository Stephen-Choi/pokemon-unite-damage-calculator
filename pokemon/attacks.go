package pokemon

// AttackOption is an enum for the different types of attacks a pokemon can do
type AttackOption string

const (
	move1       AttackOption = "move1"
	move2       AttackOption = "move2"
	uniteMove   AttackOption = "uniteMove"
	basicAttack AttackOption = "basicAttack"
)

// StatusConditions is an enum for the different types of status conditions a pokemon can inflict
type StatusConditions string

const (
	burned StatusConditions = "burned"
)

// CoolDowns is a struct containing the cooldowns of a pokemon's attacks
type CoolDowns struct {
	move1CoolDown       float64
	move2CoolDown       float64
	uniteMoveCoolDown   float64
	basicAttackCoolDown float64
}
