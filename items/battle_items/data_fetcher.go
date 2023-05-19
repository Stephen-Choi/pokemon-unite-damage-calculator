package battleitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	battleItemDataDirectory = "/Users/stephenchoi/Desktop/Projects/pokemon-unite-damage-calculator/data/battle_items"
)

// FetchBattleItemData fetches the data for a battle item
func FetchBattleItemData(battleItemName string) (battleItemData BattleItemData, err error) {
	battleItemFileName := battleItemName + ".json"

	// Read file
	battleItemFilePath := filepath.Join(battleItemDataDirectory, battleItemFileName)
	data, err := ioutil.ReadFile(battleItemFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(data, &battleItemData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	return battleItemData, nil
}
