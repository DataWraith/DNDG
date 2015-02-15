package main

// DescriptionFunc specifies a function that, given the current gamestate,
// returns (part of) the description of a given room.
type DescriptionFunc func(g Gamestate) string
