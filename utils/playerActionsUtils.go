package utils

import (
	defs "black-strat/definitions"
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

func GenerateCorrectPlayerAction(userAnswer string, playerHand *defs.PlayerHand, dealerHand *defs.DealerHand, actions *Actions) string {

	return "split"
}

func GeneratePlayerActionsKey(playerHand *defs.PlayerHand, dealerHand *defs.DealerHand) string {
	if playerHand.Card1.Symbol != playerHand.Card2.Symbol &&
		playerHand.Card1.Symbol != "A" &&
		playerHand.Card2.Symbol != "A" {
		return fmt.Sprintf("%d;%s", playerHand.Card1.Value+playerHand.Card2.Value, dealerHand.ShowingCard.Symbol)
	}
	return fmt.Sprintf("%s-%s;%s", playerHand.Card1.Symbol, playerHand.Card2.Symbol, dealerHand.ShowingCard.Symbol)
}
