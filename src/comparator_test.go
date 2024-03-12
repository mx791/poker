package main

import (
	"testing"
)

func TestPairs(t *testing.T) {
	pairA := []Card{MakeCard("Queen of Spades"), MakeCard("Queen of Hearth")}
	pairB := []Card{MakeCard("5 of Spades"), MakeCard("8 of Hearth")}

	if CompareHands(pairA, pairB) != 1 {
		t.Fatalf("Undetected pair")
	}

	pairB = []Card{MakeCard("5 of Spades"), MakeCard("5 of Hearth")}
	if CompareHands(pairA, pairB) != 1 {
		t.Fatalf("Error with pairs values")
	}

	pairA = []Card{MakeCard("8 of Spades"), MakeCard("8 of Hearth")}
	if CompareHands(pairB, pairA) != -1 {
		t.Fatalf("Error with pairs values")
	}

	pairB = []Card{MakeCard("8 of Spades"), MakeCard("8 of Hearth")}
	if CompareHands(pairB, pairA) != 0 {
		t.Fatalf("Error with pairs equality")
	}
}

func TestHighestCard(t *testing.T) {
	pairA := []Card{MakeCard("Queen of Spades"), MakeCard("5 of Hearth")}
	pairB := []Card{MakeCard("5 of Spades"), MakeCard("8 of Hearth")}

	if CompareHands(pairA, pairB) != 1 {
		t.Fatalf("Error with highest card")
	}

	pairA = []Card{MakeCard("Queen of Spades"), MakeCard("5 of Hearth"), MakeCard("8 of Hearth")}
	pairB = []Card{MakeCard("5 of Spades"), MakeCard("5 of Hearth"), MakeCard("7 of Hearth")}
	if CompareHands(pairA, pairB) != -1 {
		t.Fatalf("Error with highest card & pair")
	}

	pairA = []Card{MakeCard("Queen of Spades"), MakeCard("5 of Hearth"), MakeCard("8 of Hearth")}
	pairB = []Card{MakeCard("5 of Spades"), MakeCard("6 of Hearth"), MakeCard("7 of Hearth")}
	if CompareHands(pairA, pairB) != 1 {
		t.Fatalf("Error with highest card & pair")
	}
}

func TestSuite(t *testing.T) {
	suite := []Card{MakeCard("Queen of Spades"), MakeCard("King of Hearth"), MakeCard("10 of Hearth"), MakeCard("9 of Hearth"), MakeCard("Jack of Hearth")}
	brelan := []Card{MakeCard("5 of Spades"), MakeCard("8 of Hearth"), MakeCard("8 of Diamonds"), MakeCard("8 of Diamonds"), MakeCard("6 of Spades")}

	if CompareHands(suite, brelan) != 1 {
		t.Fatalf("Error with suite vs brelan")
	}

	doublepaire := []Card{MakeCard("5 of Spades"), MakeCard("8 of Hearth"), MakeCard("8 of Diamonds"), MakeCard("5 of Diamonds"), MakeCard("6 of Spades")}
	if CompareHands(suite, doublepaire) != 1 {
		t.Fatalf("Error with suite vs double paire")
	}

	suite2 := []Card{MakeCard("5 of Spades"), MakeCard("6 of Hearth"), MakeCard("7 of Diamonds"), MakeCard("8 of Diamonds"), MakeCard("9 of Spades")}
	if CompareHands(suite, suite2) != 1 {
		t.Fatalf("Error with suite vs suite")
	}

	if CompareHands(suite2, brelan) != 1 {
		t.Fatalf("Error with suite vs brelan")
	}
}
