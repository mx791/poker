package main

import (
	"testing"
)

func relativeErr(value float64, target float64) float64 {
	err := (value - target) / target
	if err < 0 {
		err = -err
	}
	return err
}

func TestPairValue(t *testing.T) {
	myHand := []Card{MakeCard("3 of Spades"), MakeCard("3 of Hearth")}
	value := EvalGameState(myHand, []Card{})

	if relativeErr(value, 0.6) > 0.1 {
		t.Fatalf(`Recorded %f`, value)
	}

	myHand = []Card{MakeCard("As of Spades"), MakeCard("As of Hearth")}
	value = EvalGameState(myHand, []Card{})

	if relativeErr(value, 0.9) > 0.1 {
		t.Fatalf(`Recorded %f`, value)
	}
}
