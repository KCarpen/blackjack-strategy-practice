package utils

import defs "black-strat/definitions"

type Actions = map[string]string

func GeneratePlayerActionsMap() Actions {
	actions := Actions{}

	return actions
}

func GenerateCorrectPlayerAction(userAnswer string, playerHand *defs.PlayerHand, dealerHand *defs.DealerHand, actions *Actions) string {

	return "split"
}
