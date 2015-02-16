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
+-----+   +--+--+
| 004 |---| 000 |
+-----+   +--+--+
             |
          +--+--+
          | 002 |
          +-----+
           |  ^
           |  |
           +--+
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

						g.CurrentRoom = 4
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
				Command: []string{"examine wall", "x wall", "examine stone wall", "x stone wall"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
A fairly smooth and featureless wall made out of stone bricks.
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
					if !g.HasFlag("gate-key") {
						fmt.Println(strings.TrimSpace(`
You need a key to unlock the gate.
					`))
						return false
					}

					fmt.Println(strings.TrimSpace(`
You take the key from your pocket and insert it into the lock of the gate. It fits!
You turn the key. With a loud click, the gate unlocks.
					`))
					g.SetFlag("room-000:gate-unlocked")
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
				Command: []string{"jump over gate", "jump gate", "jump wall", "jump stone wall", "vault", "vault gate", "vault wall", "vault stone wall"},
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
To the west is a smooth stone wall. The road is running in north-south
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
					if g.Rng.Float64() < 0.2 {
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

	2: Room{
		ID: 2,
		DescriptionFuncs: []DescriptionFunc{
			makeDescFunc(`
To the west is a relatively smooth stone wall. A road is running in north-south
direction. On the other side of the road is dry-land with some bushes scattered
about, but no clear destination in sight.
			`),
		},
		Actions: []Action{
			Action{
				Command: []string{"go north", "north"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You walk north until you are back at the gate.
					`))
					g.CurrentRoom = 0
					return false
				},
			},

			Action{
				Command: []string{"go south", "south"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You continue walking south, but all you can see is that stupid wall and the
deserted road.
					`))
					return false
				},
			},

			Action{
				Command: []string{"examine bushes", "x bushes", "examine plants", "x plants"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
The bushes look dry and not very healthy.
					`))
					return false
				},
			},
		},
	},

	3: Room{
		ID: 3,
		DescriptionFuncs: []DescriptionFunc{
			makeDescFunc(`
You notice a small fountain set into the wall to the west. There are some
bushes on the other side of the road.
			`),

			func(g *Gamestate) string {
				if !g.HasFlag("gate-key") {
					return "A key is lying in the fountain's basin."
				}

				return ""
			},
		},

		Actions: []Action{
			Action{
				Command: []string{"go north", "north"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You really don't want to go any further north, now that you have found
something.
					`))
					return false
				},
			},

			Action{
				Command: []string{"go south", "south"},
				Func: func(g *Gamestate) bool {
					g.CurrentRoom = 1
					return true
				},
			},

			Action{
				Command: []string{"examine fountain", "x key", "examine key", "x fountain"},
				Func: func(g *Gamestate) bool {
					if !g.HasFlag("gate-key") {
						fmt.Println(strings.TrimSpace(`
There is a key lying in the basin of the little fountain. Water is splashing
right on top of it.
					`))
						return false
					}
					fmt.Println(strings.TrimSpace(`
A lively little fountain. There is nothing special about it.
					`))
					return false
				},
			},

			Action{
				Command: []string{"examine road", "x road"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
A dust-caked asphalt road with faded lane markings.
					`))
					return false
				},
			},

			Action{
				Command: []string{"drink water"},
				Func: func(g *Gamestate) bool {
					if !g.HasFlag("room-003:not-thirsty") {
						fmt.Println(strings.TrimSpace(`
You're not sure if it is such a great idea to drink water from a random
fountain in the middle of nowhere, but you're very thirsty by now. You scoop up
the water with both hands and drink, repeatedly.
					`))
						g.SetFlag("room-003:not-thirsty")
						return false
					}

					fmt.Println(strings.TrimSpace(`
You're not really thirsty anymore.
				`))
					return false

				},
			},

			Action{
				Command: []string{"take the key", "take key"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You take the key out of the fountain's basin and shake the water off. You
pocket the key.
					`))
					g.SetFlag("gate-key")
					return false
				},
			},

			Action{
				Command: []string{"examine bushes", "x bushes", "examine plants", "x plants"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
The bushes look dried out and unhealthy. You could probably scoop up a little
water from the fountain, carry it across the road, and water the plants.
					`))
					return false
				},
			},

			Action{
				Command: []string{"water bushes", "water plants"},
				Func: func(g *Gamestate) bool {
					fmt.Println(strings.TrimSpace(`
You scoop up a handfull of water and carry it across the road. Most of the
water doesn't make it across the road, but you imagine that the bush you just
watered at least a little, is thankful and owes you something. Or something.
					`))
					return false
				},
			},
		},
	},

	4: Room{
		ID: 4,
		DescriptionFuncs: []DescriptionFunc{
			makeDescFunc(`
Behind the gate is a gravel path that leads to an oddly familiar house. Hey,
this is your house! But what is it doing here, in the middle of nowhere?

The door is slightly ajar, and you enter. You see yourself sleeping in your
bed. Suddenly you are in the bed, and as you open your eyes, you realize that
you had a rather strange dream. You wonder idly what the gate and the key
represent in your subconsciousness, and which real-life abilities or feelings
you might have locked away...

*** YOU'VE WON THE GAME ***

Enter 'exit' to quit.
			`),
		},
	},
}
