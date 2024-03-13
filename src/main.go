package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func main() {
	rand.Seed(time.Now().UnixNano())
	randomPlayer := RandomPlayer{}
	aiPlayer := ProbabilistPlayer{}

	outcome := PlayNGame([]GameBot{randomPlayer, aiPlayer, randomPlayer, randomPlayer})
	fmt.Printf("%v", outcome)
}*/

func main() {
	rand.Seed(time.Now().UnixNano())
	NUM_SIMS = 10_000
	aScore := 0.0
	randomPlayer := RandomPlayer{}
	aiPlayer := ProbabilistPlayer{}
	players := []GameBot{aiPlayer}
	for i := 0; i < 8; i++ {
		players = append(players, randomPlayer)
		for e := 0; e < 150; e++ {
			outcome := PlayNGame(players)
			aScore += outcome[0]
		}
		fmt.Printf("AI-score vs %d random agent: %f \n", i+1, aScore)
	}

}
