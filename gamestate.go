package main

import (
	"math/rand"
)

// Gamestate holds the current gamestate
type Gamestate struct {
	Rng *rand.Rand

	flags map[string]struct{}
}

// NewGamestate creates a new gamestate for the beginning of the Game
func NewGamestate(rngSeed int64) Gamestate {
	return Gamestate{
		Rng: rand.New(rand.NewSource(rngSeed)),

		flags: make(map[string]struct{}),
	}
}

// HasFlag returns whether the given flag is enabled in the Gamestate
func (g *Gamestate) HasFlag(flag string) bool {
	_, ok := g.flags[flag]
	return ok
}

// SetFlag enables the given flag in the Gamestate
func (g *Gamestate) SetFlag(flag string) {
	g.flags[flag] = struct{}{}
}

// UnsetFlag disables the given flag in the Gamestate
func (g *Gamestate) UnsetFlag(flag string) {
	delete(g.flags, flag)
}

// ToggleFlag toggles the given flag in the Gamestate
func (g *Gamestate) ToggleFlag(flag string) {
	if g.HasFlag(flag) {
		g.UnsetFlag(flag)
	} else {
		g.SetFlag(flag)
	}
}
