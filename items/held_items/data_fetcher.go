package helditems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	heldItemDataDirectory = "/Users/stephenchoi/Desktop/Projects/pokemon-unite-damage-calculator/data/held_items"
)

// FetchHeldItemData fetches the data for a held item
func FetchHeldItemData(heldItemName string) (heldItemData HeldItemData, err error) {
	heldItemFileName := heldItemName + ".json"

	// Read file
	heldItemFilePath := filepath.Join(heldItemDataDirectory, heldItemFileName)
	data, err := ioutil.ReadFile(heldItemFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(data, &heldItemData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	return heldItemData, nil
}
