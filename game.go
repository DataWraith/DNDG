package main

import (
	"fmt"
	"strings"
)

func makeDescFunc(desc string) func(g *Gamestate) string {
	d := strings.TrimSpace(desc)
	return func(g *Gamestate) string {
		return d
	}
}

/*

	+-----+
	| 001 |
	+--+--+
	   |
	+--+--+   +-----+
	| 000 |---| 003 |
	+--+--+   +-----+
	   |
	+--+--+
	| 002 |
	+-----+

*/

// Rooms holds all rooms in the game. Note that rooms can also be outdoor
// locations.
var Rooms = map[int]Room{
	0: Room{
		ID: 0,
		DescriptionFuncs: []DescriptionFunc{
			func(g *Gamestate) string {
				if !g.HasFlag("room-000:gate-open") {
					return strings.TrimSpace(`
You are standing in front of a wrought iron gate. It is fairly massive, with
spikes on top. The gate is currently closed, and next to it is a stone wall
three meters in height.
					`)
				}

				return strings.TrimSpace(`
You are standing in front of a wrought iron gate. It is fairly massive, with
spikes on top. The gate is currently open.
				`)
			},

			makeDescFunc(`
A rugged asphalt road is running beside the gate in north-south direction. Not
a single car is in sight. On the other side of the road, there is nothing of
interest. Only barren dry-land as far as the eye can see, punctuated by the
occasional bush. The bushes don't look very healthy either.
			`),
		},

		Actions: []Action{
			Action{
				Command: "go north",
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 1
					return true
				},
			},

			Action{
				Command: "go south",
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 2
					return true
				},
			},

			Action{
				Command: "go through gate",
				Func: func(g *Gamestate) bool {
					if g.HasFlag("room-000:gate-open") {
						fmt.Println(strings.TrimSpace(`
You go through the wrought-iron gate.
						`))

						g.CurrentRoom = 3
						return true
					}

					fmt.Println(strings.TrimSpace(`
The wrought-iron gate is currently closed.
					`))
					return false
				},
			},

			Action{
				Command: "examine gate",
				Func: func(g *Gamestate) bool {
					if !g.HasFlag("room-000:gate-open") {
						fmt.Println(strings.TrimSpace(`
The gate is an intricately wrought-iron structure. It is currently closed. The
bars form the initials 'DNDG', and you idly wonder whether that stands for "Do
not dare to go in" or other equally unpleasant expansions of the acronym.
						`))
						return false
					}

					fmt.Println(strings.TrimSpace(`
The gate is an intricately wrought-iron structure. It is slightly ajar. The
bars form the initials 'DNDG', and you idly wonder whether that stands for "Do
not dare to go in" or other equally unpleasant expansions of the acronym.
					`))
					return false
				},
			},

			Action{
				Command: "open gate",
				Func: func(g *Gamestate) bool {
					if g.HasFlag("room-000:gate-open") {
						fmt.Println(strings.TrimSpace(`
The gate is already open.
						`))
						return false
					}

					if !g.HasFlag("room-000:gate-unlocked") {
						fmt.Println(strings.TrimSpace(`
You depress the handle and push, but the gate does not budge. Apparently it is
locked.
						`))
						return false
					}

					fmt.Println(strings.TrimSpace(`
You depress the handle and push. The gate swings open with an ear-rending
creaking noise.
					`))
					g.SetFlag("room-000:gate-open")
					return false
				},
			},

			Action{
				Command: "close gate",
				Func: func(g *Gamestate) bool {
					if g.HasFlag("room-000:gate-open") {
						fmt.Println(strings.TrimSpace(`
You pull on the gate. With another ear-rending creak, it closes.
						`))
						g.UnsetFlag("room-000:gate-open")
						return false
					}

					fmt.Println(strings.TrimSpace(`
The gate is already closed.
					`))
					return false
				},
			},
		},
	},
}
