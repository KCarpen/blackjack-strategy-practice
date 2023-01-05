package main

import (
	"black-strat/games"
	"fmt"
)

func main() {
	switch game := selectGame(); game {
	case 0:
		games.RunBasicStrategy()
	case 1:
		games.RunCountStrategy()
	default:
		games.RunBasicStrategy()
	}
}

func selectGame() int {
	var selectedGame int

	fmt.Println("What game would you like to play? (type the number associated with the game)")
	fmt.Println("0. basic strategy practice")
	fmt.Println("1. counting cards practice")
	fmt.Print(">>> ")
	fmt.Scan(&selectedGame)

	if selectedGame != 0 && selectedGame != 1 {
		return 0
	}
	return selectedGame
}
