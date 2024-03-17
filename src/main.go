package main

/*
func main() {
	rand.Seed(time.Now().UnixNano())
	randomPlayer := RandomPlayer{}
	aiPlayer := ProbabilistPlayer{}

	outcome := PlayNGame([]GameBot{randomPlayer, aiPlayer, randomPlayer, randomPlayer})
	fmt.Printf("%v", outcome)
}*/

/*
func main() {

	rand.Seed(time.Now().UnixNano())
	NUM_SIMS = 10_000
	aScore := 0.0
	randomPlayer := RandomPlayer{0.333, 0.333, 0.333}
	// aiPlayer := ProbabilistPlayer{}
	//players := []GameBot{aiPlayer}
	players := []GameBot{randomPlayer}

	iters := 5_000.0

	for i := 0; i < 5; i++ {
		players = append(players, randomPlayer)
		for e := 0.0; e < iters; e++ {
			outcome := PlayNGame(players)
			aScore += outcome[0]
		}
		fmt.Printf("AI-score vs %d random agent: %f \n", i+1, aScore/iters)
	}
}
*/
