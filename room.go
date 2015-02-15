package main

// Room represents a single location in the game. It has a description and
// a set of actions that can be taken in the room. These actions act on the
// global Gamestate.
type Room struct {
	DescriptionFuncs []DescriptionFunc
}

// Description returns the description of a room, built from the current
// Gamestate and the room's array of DescriptionFuncs.
func (r Room) Description(g Gamestate) string {
	result := ""
	for _, p := range r.DescriptionFuncs {
		result += p(g)
		result += "\n\n"
	}
	result = result[0 : len(result)-2]
	return result
}
