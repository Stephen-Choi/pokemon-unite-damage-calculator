package main

import (
	"fmt"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
)

func main() {
	pikachu, err := pokemon.NewPikachu(13, []items.HeldItemName{items.AttackWeight, items.BuddyBarrier, items.ChoiceSpecs}, items.FluffyTail)
	if err != nil {
		fmt.Println("Error creating Pikachu:", err)
		return
	}

	fmt.Println(pikachu)
}
