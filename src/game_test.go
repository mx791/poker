package main

import (
	"testing"
	"fmt"
)

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	for i:=0; i<150; i++ {
		if i%2==0 {
			aReward += PlayGame(pA, pB)
		} else {
			aReward += -PlayGame(pB, pA)
		}
	}
	fmt.Printf("Random vs Random %f", aReward)
}
