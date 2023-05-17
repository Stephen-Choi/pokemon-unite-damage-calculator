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
