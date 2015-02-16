package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/GeertJohan/go.linenoise"
)

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

			func(g *Gamestate) string {
				return strings.TrimSpace(`
A rugged asphalt road is running beside the gate in north-south direction. Not
a single car is in sight. On the other side of the road, there is nothing of
interest. Only barren dry-land as far as the eye can see, punctuated by the
occasional bush. The bushes don't look very healthy either.
				`)
			},
		},
	},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGamestate(rand.Int63())

	linenoise.AddHistory("help")

	for {
		if _, ok := Rooms[g.CurrentRoom]; !ok {
			log.Fatalf("transitioned to undefined Room #%3d", g.CurrentRoom)
		}

		// Print the current Room's description
		fmt.Print(Rooms[g.CurrentRoom].Description(g))

		// Get the user's command
		line, err := linenoise.Line("> ")
		fmt.Println()

		// Exit the game if the user wants to leave
		if strings.EqualFold(line, "exit") || strings.EqualFold(line, "quit") || err == linenoise.KillSignalError {
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		err = linenoise.AddHistory(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}
