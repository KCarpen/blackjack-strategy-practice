package main

import (
	"black-strat/utils"
	"fmt"
)

func main() {
	outcomes := utils.GeneratePlayerActionsMap()

	for {
		playerHand, dealerHand := utils.GenerateHands()
		fmt.Printf("Player: %s, %s\n", playerHand.Card1.Symbol, playerHand.Card2.Symbol)
		fmt.Printf("Dealer: %s\n", dealerHand.ShowingCard.Symbol)
		fmt.Println("Should you stand, hit, split, or double?")
		fmt.Print(">> ")

		var userAnswer string
		fmt.Scan(&userAnswer)

		correctPlayerAction := utils.GenerateCorrectPlayerAction(userAnswer, &playerHand, &dealerHand, &outcomes)

		if userAnswer == correctPlayerAction {
			fmt.Print("Correct!\n\n")
		} else {
			fmt.Printf("Incorrect! You should %s\n\n", correctPlayerAction)
		}
	}
}
