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
	| 003 |
	+-----+
           |
	   x
           |
	+-----+
	| 001 |
	+--+--+
	   |
	+--+--+   +-----+
	| 000 |---| 004 |
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
spikes on top. The gate is currently open. Next to it is a stone wall three
meters in height.
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
				Command: []string{"go north", "north"},
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 1
					return true
				},
			},

			Action{
				Command: []string{"go south", "south"},
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 2
					return true
				},
			},

			Action{
				Command: []string{"go through gate", "go gate"},
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
				Command: []string{"examine gate", "x gate"},
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
				Command: []string{"open gate", "push gate"},
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
				Command: []string{"unlock gate", "unlock", "use key", "use key on gate", "use key on lock"},
				Func: func(g *Gamestate) bool {
					// TODO: Check inventory for key
					fmt.Println(strings.TrimSpace(`
You need a key to unlock the gate.
					`))
					return false
				},
			},

			Action{
				Command: []string{"close gate", "pull gate"},
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

			Action{
				Command: []string{"climb gate", "climb wall", "climb stone wall"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
The spikes on top of the gate look quite menacing, and the smooth stone walls
on either side don't look very climbable either. You decide to try something
else.
					`))
					return false
				},
			},

			Action{
				Command: []string{"jump over", "jump gate", "jump wall", "jump stone wall", "vault", "vault gate", "vault wall", "vault stone wall"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You'd need to be an olympic athlete to jump that high. And you'd need to know
how to pole vault. And you'd need a pole. And you'd need to not impale yourself
on the gate's spikes...
					`))
					return false
				},
			},

			Action{
				Command: []string{"examine road", "x road"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
It's an asphalt road. The lane markings are old and faded out, and dust covers
much of the surface. You notice ants scurring around some of the cracks in the surface.
					`))
					g.SetFlag("room-000:found-ants")
					return false
				},
			},

			Action{
				Command: []string{"examine ants", "x ants"},
				Func: func(g *Gamestate) bool {
					if !g.HasFlag("room-000:found-ants") {
						fmt.Println(strings.TrimSpace(`
There are no ants here.
						`))
						return false
					}

					fmt.Println(strings.TrimSpace(`
Little reddish-brown ants are scurrying around on the surface of the road. You
could stomp on them, but they're probably too tiny to actually get squished by
that. Besides, what have they done to you, to deserve such a fate?
					`))
					return false
				},
			},

			Action{
				Command: []string{"stomp ants"},
				Func: func(g *Gamestate) bool {
					if !g.HasFlag("room-000:found-ants") {
						fmt.Println("There are no ants here.")
						return false
					}

					fmt.Println(strings.TrimSpace(`
Seriously? Oh, okay...

You try to stomp the ants, but nothing satisfying happens.
					`))
					return false
				},
			},

			Action{
				Command: []string{"examine bushes", "x bushes", "examine plants", "x plants", "examine leaves", "x leaves"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You cross the road to look at the nearest bush. The leaves look dry. The bush
seems to need watering.
					`))
					return false
				},
			},

			Action{
				Command: []string{"water bushes", "water plants"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
Very funny. You don't even have any water to quench your own thirst.
					`))
					return false
				},
			},
		},
	},

	1: Room{
		ID: 1,
		DescriptionFuncs: []DescriptionFunc{
			makeDescFunc(`
To your left is a smooth stone wall. The road is running in north-south
direction to your right.
			`),

			func(g *Gamestate) string {
				if g.Rng.Float64() < 0.33 {
					return strings.TrimSpace(`
Despite its overall smoothness, some of the bricks in the wall appear to be
cracked.
					`)
				}
				return ""
			},

			func(g *Gamestate) string {
				if g.Rng.Float64() < 0.20 {
					return strings.TrimSpace(`
A small rodent scurries along the base of the wall.
					`)
				}
				return ""
			},
		},

		Actions: []Action{
			Action{
				Command: []string{"go north", "north"},
				Func: func(g *Gamestate) bool {
					if g.Rng.Float64() < 0.1 {
						g.CurrentRoom = 3
					}
					return true
				},
			},

			Action{
				Command: []string{"go south", "south"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You walk south until you are back at the gate.
					`))
					g.CurrentRoom = 0
					return false
				},
			},

			Action{
				Command: []string{"examine bricks", "x bricks", "examine cracks", "x cracks", "examine cracked bricks", "x cracked bricks"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
There is nothing remarkable about those cracked bricks.
					`))
					return false
				},
			},

			Action{
				Command: []string{"examine rodent", "x rodent", "examine small rodent", "x small rodent", "examine mouse", "x mouse"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
It seems to have vanished from sight.
					`))
					return false
				},
			},
		},
	},
}
