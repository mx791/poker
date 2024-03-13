package main

import (
	"fmt"
	"math/rand"
	"strings"
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

func MakeCard(name string) Card {
	value := 0
	family := 0
	for i := 0; i < len(values); i++ {
		if strings.Contains(name, values[i]) {
			value = i
			break
		}
	}
	for i := 0; i < len(families); i++ {
		if strings.Contains(name, families[i]) {
			family = i
			break
		}
	}
	return Card{value, family}
}

type CardGenerator struct {
	passed map[string]bool
}

func MakeCardGenerator() CardGenerator {
	return CardGenerator{make(map[string]bool)}
}

func (c1 Card) Equals(c2 Card) bool {
	return c1.ToString() == c2.ToString()
}

func genCard() Card {
	return Card{rand.Intn(13), rand.Intn(4)}
}

func (gen *CardGenerator) Next() Card {
	card := genCard()
	cardFound, ok := gen.passed[card.ToString()]
	for cardFound && ok {
		card = genCard()
		cardFound, ok = gen.passed[card.ToString()]
	}
	gen.passed[card.ToString()] = true
	return card
}

func (gen *CardGenerator) Remove(card Card) {
	gen.passed[card.ToString()] = false
}
