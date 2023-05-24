package helditems

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetHeldItem(t *testing.T) {
	t.Run("success retrieving all held items", func(t *testing.T) {
		for _, heldItemName := range playableHeldItems {
			_, err := GetHeldItem(heldItemName, lo.ToPtr(0))
			assert.NoError(t, err, fmt.Sprintf("failed fetching item: %s", heldItemName))
		}
	})
}
