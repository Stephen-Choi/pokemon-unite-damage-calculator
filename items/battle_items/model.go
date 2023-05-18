package battleitems

const (
	FluffyTailName = "fluffyTail"
	XAttackName    = "xAttack"
)

var playableBattleItems = []string{
	FluffyTailName,
	XAttackName,
}

// IsBattleItemPlayable checks if the given battle item exists in game
func IsBattleItemPlayable(battleItemName string) bool {
	for _, playableBattleItem := range playableBattleItems {
		if battleItemName == playableBattleItem {
			return true
		}
	}
	return false
}
