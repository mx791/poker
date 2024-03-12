package main

import (
	"testing"
	"fmt"
)

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	for i:=0; i<500; i++ {
		if i%2==0 {
			aReward += PlayGame(pA, pB)
		} else {
			aReward += -PlayGame(pB, pA)
		}
	}
	fmt.Printf("Random vs Random %f", aReward)
}

func TestRandomvsproba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	NUM_SIMS  = 10_000
	for i:=0; i<500; i++ {
		if i%2==0 {
			aReward += PlayGame(pA, pB)
		} else {
			aReward += -PlayGame(pB, pA)
		}
	}
	fmt.Printf("Probabilist vs Random %f", aReward)
}

func TestProba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := ProbabilistPlayer{}
	aReward := 0.0
	NUM_SIMS  = 10_000
	for i:=0; i<500; i++ {
		if i%2==0 {
			aReward += PlayGame(pA, pB)
		} else {
			aReward += -PlayGame(pB, pA)
		}
	}
	fmt.Printf("Probabilist vs Probabilist %f", aReward)
}
