package main

import (
  "math/rand"
)

const (
  ACTION_FOLLOW = 1
  ACTION_RAISE = 2
  ACTION_SLEEP = 3
)

const GameBot interface {
  PlayFirst(myCards []Card, communCards []Card, totalPotValue float64) int
  PlayNormal(myCards []Card, communCards []Card, totalPotValue float64, betValue float64) int
  ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, taregtValue float64) bool
}

func PlayGame(playerA GameBot, playerB GameBot) float64 {
  
  gen := MakeCardGenerator()
  cardA := []Card{gen.Next(), gen.Next()}
  cardB := []Card{gen.Next(), gen.Next()}
  pot := 0.0
  currentRaise := 0.0
  communCards := make([]Card, 0)

  for id, val := range []int{0, 3, 1, 1} {
    for i:=0; i<val; i++ {
        communCards = append(communCards, gen.Next())
    }
    players := []GameBot{playerA, playerB}
    playersCards := [][]Card{cardA, cardB}
    if id % 2 == 1 {
      players = []GameBot{playerB, playerA}
      playersCards = [][]Card{cardB, cardA}
    }
    firstEngaged := 0.0
    currentPot := 0.0
    betValue := 0.0
    firstAction := players[0].PlayFirst(playersCards[0], communCards, pot)
    if firstAction == ACTION_FOLLOW {
      return 0
    } else if firstAction == ACTION_RAISE {
      betValue += 1.0
      firstEngaged = 1.0
      currentPot += betValue
    } else if firstAction == ACTION_SLEEP {
      if id % 2 == 0 {
        return -pot
      } else {
        return pot
      }
    }
    secondAction := players[1].PlayNormal(playersCards[1], communCards, pot, betValue)
    if secondAction == ACTION_SLEEP {
      pot += currentPot
      if id % 2 == 1 {
        return -pot
      } else {
        return pot
      }
    } else if secondAction == ACTION_FOLLOW {
      currentPot += betValue
    } else if secondAction == ACTION_RAISE {
      betValue += 1
      if players[0].ShouldFollow(playersCards[0], communCards, pot, firstEngaged, betValue) {
        currentPot += betValue - firstEngaged
      } else {
        pot += currentPot
        if id % 2 == 0 {
          return -pot
        } else {
          return pot
        }
      }
    }
    pot += currentPot
  }

  // endgame
  winner := CompareHands(append(cardA, communCards...), append(cardB, communCards...))
  if winner == 0 {
    return 0.0
  } else if winner == 1 {
    return pot
  }
  return -pot
}
