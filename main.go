package main

import (
	"black-strat/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type PlayerHand struct {
	card1 Card
	card2 Card
}

type DealerHand struct {
	showingCard Card
}

type Card struct {
	symbol string
	value  int
}

func main() {
	outcomes := utils.GenerateOutcomes()

	for {
		playerHand, dealerHand := generateHands()
		fmt.Printf("Player: %s, %s\n", playerHand.card1.symbol, playerHand.card2.symbol)
		fmt.Printf("Dealer: %s\n", dealerHand.showingCard.symbol)
		fmt.Println("Should you stand, hit, split, or double?")
		fmt.Print(">> ")

		var userAnswer string
		fmt.Scan(&userAnswer)

		correctAction := generateCorrectAction(userAnswer, &playerHand, &dealerHand, &outcomes)

		if userAnswer == correctAction {
			fmt.Print("Correct!\n\n")
		} else {
			fmt.Printf("Incorrect! You should %s\n\n", correctAction)
		}
	}
}

func generateHands() (PlayerHand, DealerHand) {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	var randomCards [3]Card

	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(len(cards))

		randomCard := Card{symbol: cards[randomInt], value: computeCardValue(cards[randomInt])}
		randomCards[i] = randomCard
	}

	lowerPlayerCard, higherPlayerCard := sortCardsDesc(randomCards[0], randomCards[2], cards)

	playerHand := PlayerHand{card1: lowerPlayerCard, card2: higherPlayerCard}
	dealerHand := DealerHand{showingCard: randomCards[1]}

	return playerHand, dealerHand
}

func computeCardValue(symbol string) int {
	if symbol == "A" {
		return 0
	}
	if symbol == "J" || symbol == "Q" || symbol == "K" {
		return 10
	}

	num, _ := strconv.Atoi(symbol)
	return num
}

func sortCardsDesc(card1 Card, card2 Card, cards []string) (Card, Card) {
	card1Idx := utils.IndexOf(cards, card1.symbol)
	card2Idx := utils.IndexOf(cards, card2.symbol)

	if card1Idx <= card2Idx {
		return card1, card2
	}
	return card2, card1
}

func generateCorrectAction(userAnswer string, playerHand *PlayerHand, dealerHand *DealerHand, outcomes *utils.Outcomes) string {

	return "split"
}
