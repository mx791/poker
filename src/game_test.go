package main

import (
	"testing"
	"fmt"
)

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	for i:=0; i<100; i++ {
		if PlayGame(pA, pB) > 0.0 {
			aReward += 1.0
		}
		if PlayGame(pB, pA) < 0.0 {
			aReward += 1.0
		}
	}
	fmt.Printf("Random vs Random %f", aReward/2.0)
}

func TestRandomvsproba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	NUM_SIMS  = 10_000
	for i:=0; i<100; i++ {
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
	for i:=0; i<100; i++ {
		if i%2==0 {
			aReward += PlayGame(pA, pB)
		} else {
			aReward += -PlayGame(pB, pA)
		}
	}
	fmt.Printf("Probabilist vs Probabilist %f", aReward)
}
