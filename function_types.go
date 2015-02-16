package main

// InitializationFunc specifies a function that is run once for a room to
// initialize it. Since the initialization may depend on the current Gamestate,
// it is called the first time a Room's Description method is called.
type InitializationFunc func(g *Gamestate)

// DescriptionFunc specifies a function that, given the current Gamestate,
// returns (part of) the description of a given room.
type DescriptionFunc func(g *Gamestate) string

// ActionFunc specifies a function that performs an action (such as "go north")
// on the current Gamestate. The return value indicates whether the room
// description should be displayed after the action is complete. For example,
// if the action changes the current room, the new room's description should be
// displayed, but if you merely pick up an item ("You pick up the scissors"),
// we do not need to display the room description again.
type ActionFunc func(g *Gamestate) bool
