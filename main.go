package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/GeertJohan/go.linenoise"
)

func transformCommand(command string) string {
	// First, we need to normalize whitespace
	fields := strings.Fields(command)
	result := strings.Join(fields, " ")

	result = strings.Replace(result, " the ", "", -1)
	result = strings.Replace(result, "look at", "examine", -1)

	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGamestate(rand.Int63())

	linenoise.AddHistory("help")
	fmt.Println()

	displayDescription := true

	for {
		if _, ok := Rooms[g.CurrentRoom]; !ok {
			log.Fatalf("transitioned to undefined Room #%03d", g.CurrentRoom)
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

		// Transform the input to catch slightly different ways of phrasing a command
		line = transformCommand(line)

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
