package main

import (
	"fmt"
	helditems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
)

func main() {
	//pikachu, err := pokemon.NewPikachu(13, []helditems.Name{helditems.ScoreShield}, battleitems.FluffyTail)
	//if err != nil {
	//	fmt.Println("Error creating Pikachu:", err)
	//	return
	//}
	//
	//fmt.Println(pikachu)

	heldItem, err := helditems.FetchHeldItemData(helditems.SpecialAttackSpecsName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", heldItem)
}
