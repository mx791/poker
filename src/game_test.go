package main

import (
	"math/rand"
	"testing"
	"fmt"
)

func TestRandomGame(t *testing.T) {
	pA := RandomPlayer{}
	pB := RandomPlayer{}
	aReward := 0.0
	for i:=0; i<150; i++ {
		aReward += PlayGame(pA, pB)
	}
	fmt.Printf("Random vs Random %f", aReward)
}
