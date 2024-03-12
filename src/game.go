package main

const (
  ACTION_FOLLOW = 1
  ACTION_RAISE = 2
  ACTION_SLEEP = 3
  ACTION_CALL = 4
)

const GameBot interface {
  play(myCards Card[], communCards Card[], isFirst bool, totalPotValue float64, currentPotValue float64) int
}
