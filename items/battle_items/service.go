package battleitems

import "fmt"

// GetBattleItem fetches a battle item
func GetBattleItem(battleItemName string) (battleItem BattleItem, err error) {
	if !IsBattleItemPlayable(battleItemName) {
		err = fmt.Errorf("battle item %s is not playable", battleItemName)
		return
	}

	switch battleItemName {
	case FluffyTailName:
		battleItem, err = NewFluffyTail()
	case XAttackName:
		battleItem, err = NewXAttack()
	}
	return
}