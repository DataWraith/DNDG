package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/GeertJohan/go.linenoise"
)

// DEBUG enables or disables debugging functionality for the program
const DEBUG = true

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

		switch line {
		case "help":
			fmt.Printf("Blah blah, Tab completion, blah blah\n\n")
			fmt.Printf("You can currently issue the following commands:\n\n")
			for _, action := range Rooms[g.CurrentRoom].Actions {
				fmt.Printf("* %s\n", action.Command)
			}
			fmt.Println()
			displayDescription = false
			continue

		case "i", "inv", "inventory":
			fmt.Println("Inventory is not yet implemented")

		case "flags":
			if !DEBUG {
				fmt.Println("This command is only available in DEBUG mode")
				continue
			}

			flags := make([]string, 0, len(g.flags))
			for k := range g.flags {
				flags = append(flags, k)
			}
			sort.Strings(flags)

			for _, flag := range flags {
				fmt.Printf("* %s\n", flag)
			}

		default:
			// Transform the input to catch slightly different ways of phrasing a command
			tline := transformCommand(line)

			if DEBUG && tline != line {
				log.Printf("Transformed input: %q\n\n", tline)
			}

			foundAction := false
			for _, action := range Rooms[g.CurrentRoom].Actions {
				if action.Command == tline {
					foundAction = true
					displayDescription = Rooms[g.CurrentRoom].ExecuteAction(tline, g)
					break
				}
			}

			if !foundAction {
				fmt.Printf("Sorry, I did not understand the command %q\n", line)
			}
		}

	}
}
