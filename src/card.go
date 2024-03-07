package main

import (
	"fmt"
	"math/rand/v2"
)

var families []string = []string{"Hearth", "Diamonds", "Clubs", "Spades"}
var values []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "As"}

type Card struct {
	value  int
	family int
}

func (c Card) ToString() string {
	return fmt.Sprintf("%s of %s", values[c.value], families[c.family])
}

type CardGenerator struct {
	passed map[string]bool
}

func MakeCardGenerator() CardGenerator {
	return CardGenerator{make(map[string]bool)}
}

func (c1 Card) Equals(c2 Card) bool {
	return c1.toString() == c2.toString()
}

func genCard() Card {
	return Card{rand.IntN(13), rand.IntN(4)}
}

func (gen *CardGenerator) Next() Card {
	card := genCard()
	card, found := gen.passed[card.toString()]
	for found && card {
		card = genCard()
		_, found = gen.passed[card.toString()]
	}
	gen.passed[card.toString()] = true
	return card
}

func (gen *CardGenerator) Remove(card Card) {
	gen.passed[card.toString()] = false
}
