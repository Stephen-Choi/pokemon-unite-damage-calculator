package battleitems

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetBattleItem(t *testing.T) {
	t.Run("success retrieving all battle items", func(t *testing.T) {
		for _, battleItemName := range playableBattleItems {
			_, err := GetBattleItem(battleItemName)
			assert.NoError(t, err, fmt.Sprintf("failed fetching item: %s", battleItemName))
		}
	})
}
