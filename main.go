package main

import (
	"black-strat/utils"
	"flag"
	"fmt"
)

func main() {
	csvFilename := flag.String("csv", "playerActions.csv", "a CSV file in the format 'players hand,dealers hand,recommended action'")

	flag.Parse()

	playerActionsMap := utils.GeneratePlayerActionsMap(*csvFilename)

	for {
		playerHand, dealerHand := utils.GenerateHands()
		fmt.Printf("Player: %s, %s\n", playerHand.Card1.Symbol, playerHand.Card2.Symbol)
		fmt.Printf("Dealer: %s\n", dealerHand.ShowingCard.Symbol)
		fmt.Println("Should you stand, hit, split, or double?")
		fmt.Print(">> ")

		var userAnswer string
		fmt.Scan(&userAnswer)

		correctPlayerAction := utils.GenerateCorrectPlayerAction(userAnswer, &playerHand, &dealerHand, &playerActionsMap)

		if userAnswer == correctPlayerAction {
			fmt.Print("Correct!\n\n")
		} else {
			fmt.Printf("Incorrect! You should %s\n\n", correctPlayerAction)
		}
	}
}
