package main

import (
	"math/rand"
)

const (
  ACTION_FOLLOW = 1
  ACTION_RAISE = 2
  ACTION_SLEEP = 3
)

type GameBot interface {
  Play(myCards []Card, communCards []Card, totalPotValue float64, betValue float64) int
  ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, targetValue float64) bool
}

func PlayNGame(players []GameBot) []float64 {
	gen := MakeCardGenerator()
	playerCards := make([][]Card, len(players))
	ttInvestedByPlayer := make([]float64, len(players))
	outPlayers := make([]bool, len(players))
	activesPlayers := len(players)
	communCards := make([]Card, 0)
	pot := 0.0
	playerWin := -1
	for id, val := range []int{0, 3, 1, 1} {
		currentBet := 0.0
		currentInvestedByPlayer := make([]float64, len(players))
    		for i:=0; i<val; i++ {
        		communCards = append(communCards, gen.Next())
		}
		for pId:=0; pId<len(players); pId++ {
			playerPosId := (id+pId) % len(players)
			if outPlayers[playerPosId] || activesPlayers == 1 {
				continue
			}
			action := players[playerPosId].Play(playerCards[playerPosId], communCards, pot, currentBet)
			if action == ACTION_FOLLOW {
				currentInvestedByPlayer[playerPosId] += currentBet
			} else if action == ACTION_RAISE {
				currentBet += 1
				currentInvestedByPlayer[playerPosId] += currentBet
			} else if action == ACTION_SLEEP {
				outPlayers[playerPosId] = true
				activesPlayers -= 1
			}
		}
		for pId:=0; pId<len(players); pId++ {
			playerPosId := (id+pId) % len(players)
			if outPlayers[playerPosId] || currentInvestedByPlayer[playerPosId] == currentBet || activesPlayers == 1 {
				continue
			}
			doFollow := players[playerPosId].ShouldFollow(playerCards[playerPosId], communCards, pot, currentInvestedByPlayer[playerPosId], currentBet)
			if doFollow {
				currentInvestedByPlayer[playerPosId] = currentBet
			} else {
				outPlayers[playerPosId] = true
				activesPlayers -= 1
			}
		}
		for i, val := range currentInvestedByPlayer {
			ttInvestedByPlayer[i] += val
			pot += val
		}
		if activesPlayers == 1 {
			for pId:=0; pId<len(players); pId++ {
				if !outPlayers[(id+pId) % len(players)] {
					playerWin = (id+pId) % len(players)
				}
			}
		}
    	}
	if playerWin == -1 {
		victoryCount := make([]int, len(players))
		for i:=0; i<len(players)-1; i++ {
			for e:=i+1; e<len(players); e++ {
				out :=  CompareHands(append(playerCards[i], communCards...), append(playerCards[e], communCards...))
				if out >= 0 {
					victoryCount[i] += 1
				} else {
					victoryCount[e] += 1
				}
			}
			if victoryCount[i] == len(players)-1 {
				playerWin = i
			}
 		}
	}
	payOff := make([]float64, len(players))
	for id, val := range ttInvestedByPlayer {
		payOff[id] = -val
		if id == playerWin {
			payOff[id] = pot
		}
	}
	return payOff
}

func PlayGame(playerA GameBot, playerB GameBot) float64 {
	out := PlayNGame([]GameBot{playerA, playerB})
	return out[0]
}

type RandomPlayer struct {}

func (p RandomPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64) int {
  val := rand.Intn(5)
  if val == 0 || val == 1 || val == 2 {
    return ACTION_FOLLOW
  } else if val == 3 {
    return ACTION_RAISE
  }
  return ACTION_SLEEP
}

func (p RandomPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, taregtValue float64) bool {
	return rand.Intn(5) > 2
}

type ProbabilistPlayer struct {}

func (p ProbabilistPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64) int {
	proba := EvalGameState(myCards, communCards)
	if proba > 0.7 {
		return ACTION_RAISE
	}
	if proba > 0.5 {
		return ACTION_FOLLOW
	}
	return ACTION_SLEEP
}

func (p ProbabilistPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, taregtValue float64) bool {
	return EvalGameState(myCards, communCards) > 0.5
}
