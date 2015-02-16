package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/GeertJohan/go.linenoise"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGamestate(rand.Int63())

	linenoise.AddHistory("help")
	fmt.Println()

	displayDescription := true

	for {
		if _, ok := Rooms[g.CurrentRoom]; !ok {
			log.Fatalf("transitioned to undefined Room #%3d", g.CurrentRoom)
		}

		// Print the current Room's description
		if displayDescription {
			fmt.Print(Rooms[g.CurrentRoom].Description(g))
		}

		// Get the user's command
		line, err := linenoise.Line("> ")
		line = strings.ToLower(line)

		fmt.Println()

		// Exit the game if the user wants to leave
		if line == "exit" || line == "quit" || err == linenoise.KillSignalError {
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		linenoise.AddHistory(line)

		foundAction := false
		for _, action := range Rooms[g.CurrentRoom].Actions {
			if action.Command == line {
				foundAction = true
				displayDescription = Rooms[g.CurrentRoom].ExecuteAction(line, g)
				break
			}
		}

		if !foundAction {
			fmt.Printf("Sorry, I did not understand the command %q\n", line)
		}
	}
}
