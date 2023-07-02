package utils

import (
	models "black-strat/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Actions = map[string]string

func GeneratePlayerActionsMap(csvFilename string) Actions {
	f, err := os.Open(csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, _ := csvReader.ReadAll()

	actions := Actions{}

	for _, record := range records {
		key := fmt.Sprintf("%s;%s", strings.TrimSpace(record[0]), strings.TrimSpace(record[1]))
		value := strings.TrimSpace(record[2])

		actions[key] = value
	}

	return actions
}

func GenerateCorrectPlayerAction(playerHand *models.PlayerHand, dealerHand *models.DealerHand, actions Actions) string {
	key := generatePlayerActionsKey(playerHand, dealerHand)
	// fmt.Printf("key: %s\n", key)
	return actions[key]
}

func generatePlayerActionsKey(playerHand *models.PlayerHand, dealerHand *models.DealerHand) string {
	var dealerCardSymbol string

	if dealerHand.ShowingCard.Value == 10 {
		dealerCardSymbol = "10"
	} else {
		dealerCardSymbol = dealerHand.ShowingCard.Symbol
	}

	if playerHand.Card1.Value == 10 {
		playerHand.Card1.Symbol = "10"
	}

	if playerHand.Card2.Value == 10 {
		playerHand.Card2.Symbol = "10"
	}

	if playerHand.Card1.Symbol != playerHand.Card2.Symbol &&
		playerHand.Card1.Symbol != "A" &&
		playerHand.Card2.Symbol != "A" {

		playerSum := playerHand.Card1.Value + playerHand.Card2.Value
		if playerSum >= 17 {
			playerSum = 17
			dealerCardSymbol = "x"
		}
		return fmt.Sprintf("%d;%s", playerSum, dealerCardSymbol)
	}

	lowerPlayerCard, higherPlayerCard := sortCardsDesc(playerHand.Card1, playerHand.Card2, cards)

	return fmt.Sprintf("%s-%s;%s", higherPlayerCard.Symbol, lowerPlayerCard.Symbol, dealerCardSymbol)
}
