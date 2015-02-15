package main

import (
	"math/rand"
)

// Gamestate holds the current gamestate
type Gamestate struct {
	Rng *rand.Rand
}

// NewGamestate creates a new gamestate for the beginning of the Game
func NewGamestate(rngSeed int64) Gamestate {
	return Gamestate{
		Rng: rand.New(rand.NewSource(rngSeed)),
	}
}
