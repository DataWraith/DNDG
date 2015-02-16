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

// AllCommands is a global variable holding all currently valid commands
var AllCommands []Action

func transformCommand(command string) string {
	// First, we need to normalize whitespace
	fields := strings.Fields(command)
	result := strings.Join(fields, " ")

	result = strings.Replace(result, "the ", "", -1)
	result = strings.Replace(result, "look at", "examine", -1)

	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGamestate(rand.Int63())

	linenoise.SetCompletionHandler(CommandCompletion)
	linenoise.AddHistory("help")
	fmt.Println()

	displayDescription := true

	for {
		fmt.Println()

		if _, ok := Rooms[g.CurrentRoom]; !ok {
			log.Fatalf("transitioned to undefined Room #%03d", g.CurrentRoom)
		}

		// Setup AllCommands (needed for tab-completion)
		AllCommands = Rooms[g.CurrentRoom].Actions

		// Print the current Room's description
		if displayDescription {
			fmt.Print(Rooms[g.CurrentRoom].Description(g))
		}
		displayDescription = false

		// Get the user's command
		line, err := linenoise.Line("> ")
		line = strings.ToLower(line)

		// Exit the game if the user wants to leave
		if line == "exit" || line == "quit" || err == linenoise.KillSignalError {
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		linenoise.AddHistory(line)

		switch line {
		case "":
			continue

		case "help":
			fmt.Println(strings.TrimSpace(`
Blah blah blah, Tab completion, blah blah.

You can type in 'inventory' (abbreviated as 'i') to examine your inventory.

You can type in 'commands' (abbreviated as 'c') to display a list of commands
you can use in your current location. This can contain SPOILERS, so only use
this if you are truly stuck.
			`))
			continue

		case "c", "commands":
			fmt.Printf("You can currently issue the following commands:\n\n")
			for _, action := range Rooms[g.CurrentRoom].Actions {
				fmt.Printf("* %s\n", action.Command[0])
			}
			continue

		case "i", "inventory":
			fmt.Println("Inventory is not yet implemented")

		case "f", "flags":
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

		case "xyzzy":
			fmt.Println("Nothing happens.")

		default:
			// Transform the input to catch slightly different ways of phrasing a command
			tline := transformCommand(line)

			if DEBUG && tline != line {
				log.Printf("Transformed input: %q\n\n", tline)
			}

			foundAction := false
			for _, action := range Rooms[g.CurrentRoom].Actions {
				for _, cmd := range action.Command {
					if cmd == tline {
						foundAction = true
						fmt.Println()
						displayDescription = Rooms[g.CurrentRoom].ExecuteAction(tline, g)
						break
					}
				}
			}

			if !foundAction {
				completion := CommandCompletion(tline)
				if len(completion) == 1 {
					fmt.Println(">", completion[0])
					fmt.Println()
					displayDescription = Rooms[g.CurrentRoom].ExecuteAction(completion[0], g)
					continue
				} else if len(completion) > 1 {
					fmt.Printf("\nThe command %q was ambiguous. Which of the following did you mean?\n\n", line)
					for _, c := range completion {
						fmt.Printf("* %s\n", c)
					}
					continue
				}

				fmt.Printf("Sorry, I did not understand the command %q\n", line)
			}
		}

	}
}
