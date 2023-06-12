# pokemon-unite-damage-calculator (WIP)

## Motivation:
Main goal of this project is to create a damage calculator for the game 'Pokemon Unite' to determine how long it takes for a team of pokemon to defeat objectives. The game does not provide a convenient way to test this out apart from going into practice mode and playing out entire games. This tool should help players create better strategies and understandings around when and how to rip down objectives.

## Calculations
The pokemon's levels, attacks, buffs gained and inflicted debuffs, passives, held items and battle items will all accounted for in the calculation.

The math involved in this project is from data avilable on unite-db as well as findings from the 'pokemon unite math discord'. They provided the pokemon stats, attack multipliers, how attack speed works, and many others. Big thanks to them for their work obtaining that data!

## Current project status:

The engine to simulate a battle against an objective pokemon is complete. You can find that in the ```damagecalculator``` package, specifically the ```CalculateRip()``` method. Currently the setup for pokemons: 1. Pikachu and 2. Slowbro have been implemented along with all damage related held and battle item effects. I will be adding the remaining pokemon over time.
