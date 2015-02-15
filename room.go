package main

import (
	"fmt"
)

// Room represents a single location in the game. It has a description and
// a set of actions that can be taken in the room. These actions act on the
// global Gamestate.
type Room struct {
	ID                 int
	InitializationFunc InitializationFunc
	DescriptionFuncs   []DescriptionFunc
}

// Description returns the description of a room, built from the current
// Gamestate and the room's array of DescriptionFuncs.
func (r Room) Description(g *Gamestate) string {
	// Call the initialization function of the room if it wasn't called before
	initFlag := fmt.Sprintf("room-%3d:initialized", r.ID)
	if !g.HasFlag(initFlag) {
		if r.InitializationFunc != nil {
			r.InitializationFunc(g)
		}

		g.SetFlag(initFlag)
	}

	// Build the room description
	result := ""
	for _, dfunc := range r.DescriptionFuncs {
		desc := dfunc(g)
		if desc == "" {
			continue
		}

		result += desc
		result += "\n\n"
	}
	return result
}
