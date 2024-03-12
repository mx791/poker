package main

import (
	"testing"
	"fmt"
)

func comparePlayers(pA GameBot, pB GameBot) (float64, float64) {
	aReward := 0.0
	bReward := 0.0
	for i:=0; i<100; i++ {
		if PlayGame(pA, pB) > 0.0 {
			aReward += 1.0
		} else {
			bReward += 1.0
		}
		if PlayGame(pB, pA) < 0.0 {
			aReward += 1.0
		} else {
			bReward += 1.0
		}
	}
	return aReward / 200.0, bReward/200.0
}

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	resA, resB := comparePlayers(pA, pB)
	fmt.Printf("Random vs Random %f - %f", resA, resB)
}

func TestRandomvsproba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := RandomPlayer{}
	NUM_SIMS  = 10_000
	resA, resB := comparePlayers(pA, pB)
	fmt.Printf("Random vs Random %f - %f", resA, resB)
}

func TestProba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := ProbabilistPlayer{}
	NUM_SIMS  = 10_000
	resA, resB := comparePlayers(pA, pB)
	fmt.Printf("Random vs Random %f - %f", resA, resB)
}
