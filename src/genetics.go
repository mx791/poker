package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BenchMark(bot GameBot) float64 {
	iters := 20_000.0
	randomPlayer := RandomPlayer{0.333, 0.333, 0.333, 0.5}
	players := []GameBot{bot, randomPlayer, randomPlayer, randomPlayer, randomPlayer}
	aScore := 0.0
	for e := 0.0; e < iters; e++ {
		outcome := PlayNGame(players)
		aScore += outcome[0]
	}
	return aScore / iters
}

func GeneticEvol(nPlayers int) {
	baseBot := RandomPlayer{0.333, 0.333, 0.333, 0.5}
	players := []GameBot{baseBot}
	for len(players) < nPlayers {
		players = append(players, baseBot)
	}
	for i := 0; i < 100; i++ {
		results := make([]float64, nPlayers)
		for e := 0; e < 20_000; e++ {
			res := PlayNGame(players)
			for id, v := range res {
				results[id] = results[id] + v
			}
		}
		max := results[0]
		maxId := 0
		for id, val := range results {
			if val > max {
				max = val
				maxId = id
			}
		}
		players[0] = players[maxId] //
		bestPlayer := players[maxId]
		for e := 1; e < nPlayers; e++ {
			players[e] = bestPlayer.Mutate()

		}
		fmt.Printf("%f \n", BenchMark(bestPlayer))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	GeneticEvol(5)
}
