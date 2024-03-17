package main

import (
	"math/rand"
)

type GameBot interface {
	Play(myCards []Card, communCards []Card, totalPotValue float64, betValue float64) int
	ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, targetValue float64) bool
}

type RandomPlayer struct {
	FollowProba       float64
	RaiseProba        float64
	SleepProba        float64
	SleepOnRaiseProba float64
}

type MutableAgent interface {
	GameBot
	Mutate() GameBot
}

func (p RandomPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64, potValue float64) int {
	val := rand.Float64()
	if val < p.FollowProba {
		return ACTION_FOLLOW
	} else if val > p.FollowProba && val < p.FollowProba+p.RaiseProba {
		return ACTION_RAISE
	}
	return ACTION_SLEEP
}

func (p RandomPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, taregtValue float64) bool {
	return rand.Float64() > p.SleepOnRaiseProba
}

func (p RandomPlayer) Mutate() GameBot {
	learning_rate := 0.1
	w := []float64{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}
	w = []float64{
		p.FollowProba*(1.0-learning_rate) + learning_rate*w[0],
		p.RaiseProba*(1.0-learning_rate) + learning_rate*w[1],
		p.SleepProba*(1.0-learning_rate) + learning_rate*w[2],
		p.SleepOnRaiseProba*(1.0-learning_rate) + learning_rate*w[3],
	}
	sum := w[0] + w[1] + w[2]
	w = []float64{w[0] / sum, w[1] / sum, w[2] / sum, w[3]}
	return RandomPlayer{w[0], w[1], w[2], w[3]}
}

type ProbabilistPlayer struct{}

func (p ProbabilistPlayer) Play(myCards []Card, communCards []Card, totalPotValue float64, potValue float64) int {
	proba := EvalGameState(myCards, communCards)
	if proba > 0.7 {
		return ACTION_RAISE
	}
	if proba > 0.5 {
		return ACTION_FOLLOW
	}
	return ACTION_SLEEP
}

func (p ProbabilistPlayer) ShouldFollow(myCards []Card, communCards []Card, totalPotValue float64, engagedValue float64, targetValue float64) bool {
	proba := EvalGameState(myCards, communCards)
	return (proba >= 0.5)
}
