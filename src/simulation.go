package main

var NUM_SIMS = 50_000
var DRAW_VALUE = 0.5

func EvalGameState(myCards []Card, knownCommunCards []Card) float64 {

	cumValue := 0.0
	gen := MakeCardGenerator()
	for i := 0; i < NUM_SIMS; i++ {
		otherCards := make([]Card, 5-len(knownCommunCards))
		for e := 0; e < len(otherCards); e++ {
			otherCards[e] = gen.Next()
		}
		oponent := []Card{gen.Next(), gen.Next()}
		result := CompareHands(
			append(myCards, append(knownCommunCards, otherCards...)...),
			append(oponent, append(knownCommunCards, otherCards...)...))
		if result == 1 {
			cumValue += 1.0
		}
		if result == 0 {
			cumValue += DRAW_VALUE
		}
		for _, card := range oponent {
			gen.Remove(card)
		}
		for _, card := range append(knownCommunCards, otherCards...) {
			gen.Remove(card)
		}
	}
	return cumValue / float64(NUM_SIMS)
}
