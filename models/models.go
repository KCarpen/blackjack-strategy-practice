package models

type PlayerHand struct {
	Card1 Card
	Card2 Card
}

type DealerHand struct {
	ShowingCard Card
}

type Card struct {
	Symbol string
	Value  int
}
