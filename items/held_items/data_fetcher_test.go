package helditems

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_fetchHeldItems tests the FetchHeldItemData function
func Test_fetchHeldItems(t *testing.T) {
	for _, heldItemName := range playableHeldItems {
		_, err := FetchHeldItemData(heldItemName)
		assert.NoError(t, err, fmt.Sprintf("failed fetching item: %s", heldItemName))
	}
}
