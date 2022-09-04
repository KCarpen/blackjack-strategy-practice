package utils

import (
	defs "black-strat/definitions"
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	for _, record := range records {
		fmt.Println(record)
	}

	return actions
}

func GenerateCorrectPlayerAction(userAnswer string, playerHand *defs.PlayerHand, dealerHand *defs.DealerHand, actions *Actions) string {

	return "split"
}
