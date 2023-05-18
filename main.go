package main

import (
	"fmt"
	battleitems "github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/battle_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/items/held_items"
	"github.com/Stephen-Choi/pokemon-unite-damage-calculator/pokemon"
)

func main() {
	pikachu, err := pokemon.NewPikachu(13, []helditems.Name{helditems.ScoreShield}, battleitems.FluffyTail)
	if err != nil {
		fmt.Println("Error creating Pikachu:", err)
		return
	}

	fmt.Println(pikachu)
}
