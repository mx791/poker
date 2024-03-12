package main

import (
	"testing"
	"fmt"
	"math/rand"
)


type TalkativeRandomPlayer struct {}

func (p TalkativeRandomPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64, betValue float64) int {
  val := rand.Intn(5)
  fmt.Printf("A mon tour, il y a %d cartes au milieu, la mise est à %f \n", len(communCards), betValue)
  if val == 0 || val == 1 || val == 2 {
	  fmt.Println("Je suis")
    return ACTION_FOLLOW
  } else if val == 3 {
	  fmt.Println("Je relance")
    return ACTION_RAISE
  }
	fmt.Println("Je me couche")
  return ACTION_SLEEP
}

func (p TalkativeRandomPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, targetValue float64) bool {
	fmt.Printf("On me relance")
	return rand.Intn(5) > 2
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
