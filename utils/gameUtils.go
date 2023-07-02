package utils

import (
	models "black-strat/models"
	"math/rand"
	"strconv"
	"time"
)

var cards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func GenerateHands() (models.PlayerHand, models.DealerHand) {
	var randomCards [3]models.Card

	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(len(cards))

		randomCard := models.Card{Symbol: cards[randomInt], Value: computeCardValue(cards[randomInt])}
		randomCards[i] = randomCard
	}

	lowerPlayerCard, higherPlayerCard := sortCardsDesc(randomCards[0], randomCards[2], cards)

	playerHand := models.PlayerHand{Card1: lowerPlayerCard, Card2: higherPlayerCard}
	dealerHand := models.DealerHand{ShowingCard: randomCards[1]}

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

func sortCardsDesc(card1 models.Card, card2 models.Card, cards []string) (models.Card, models.Card) {
	card1Idx := indexOf(cards, card1.Symbol)
	card2Idx := indexOf(cards, card2.Symbol)

	if card1Idx <= card2Idx {
		return card1, card2
	}
	return card2, card1
}

func indexOf(slice []string, target string) int {
	for idx, v := range slice {
		if v == target {
			return idx
		}
	}
	return -1
}
