package battleitems

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_fetchHeldItems tests the FetchHeldItemData function
func Test_fetchHeldItems(t *testing.T) {
	for _, battleItemName := range playableBattleItems {
		_, err := FetchBattleItemData(battleItemName)
		assert.NoError(t, err, fmt.Sprintf("failed fetching item: %s", battleItemName))
	}
}
