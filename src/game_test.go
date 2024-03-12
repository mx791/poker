package main

import (
	"testing"
	"fmt"
)

func comparePlayers(pA GameBot, pB GameBot) float64 {
	aReward := 0.0
	for i:=0; i<100; i++ {
		aReward += PlayGame(pA, pB)
		aReward += -PlayGame(pB, pA)
	}
	return aReward / 100.0
}

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	resA := comparePlayers(pA, pB)
	fmt.Printf("Random vs Random %f", resA)
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
	resA := comparePlayers(pA, pB)
	fmt.Printf("Random vs Random %f", resA)
}
