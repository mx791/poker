package main

import (
	"testing"
	"fmt"
	"math/rand"
)


type TalkativeRandomPlayer struct {}

func (p TalkativeRandomPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64, betValue float64) int {
  	proba := EvalGameState(myCards, communCards)
	fmt.Printf("Win_proba=%f \n", proba)
	if proba > 0.7 {
		fmt.Printf("je relance")
		return ACTION_RAISE
	}
	if proba > 0.5 {
		fmt.Printf("je suis")
		return ACTION_FOLLOW
	}
	fmt.Printf("Je me couche")
	return ACTION_SLEEP
}

func (p TalkativeRandomPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, targetValue float64) bool {
	proba := EvalGameState(myCards, communCards)
	fmt.Printf("On me relance \n Win_proba=%f \n", proba)
	return proba > 0.5
}

func comparePlayers(pA GameBot, pB GameBot) float64 {
	aReward := 0.0
	for i:=0; i<100; i++ {
		aReward += PlayGame(pA, pB)
		aReward += -PlayGame(pB, pA)
	}
	return aReward / 200.0
}

func TestGame(t *testing.T) {
	fmt.Println("Nouvelle partie")
	pA := TalkativeRandomPlayer{}
	pB := RandomPlayer{}
	out := PlayNGame([]GameBot{pA, pB})
	fmt.Printf("Issue: %v", out)
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
	resA := comparePlayers(pA, pB)
	fmt.Printf("Probabilist vs Random %f", resA)
}

func TestProba(t *testing.T) {
	pA := ProbabilistPlayer{}
	pB := ProbabilistPlayer{}
	NUM_SIMS  = 10_000
	resA := comparePlayers(pA, pB)
	fmt.Printf("Probabilist vs Probabilist %f", resA)
}
