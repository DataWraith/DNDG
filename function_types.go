package main

// InitializationFunc specifies a function that is run once for a room to
// initialize it. Since the initialization may depend on the current Gamestate,
// it is called the first time a Room's Description method is called.
type InitializationFunc func(g *Gamestate)

// DescriptionFunc specifies a function that, given the current gamestate,
// returns (part of) the description of a given room.
type DescriptionFunc func(g *Gamestate) string
