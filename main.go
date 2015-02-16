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
	result = strings.Replace(result, "over ", "", -1)
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

	if DEBUG {
		for k, v := range Rooms {
			if k != v.ID {
				log.Fatalf("ID of room %03d does not match its index, %03d", v.ID, k)
			}
		}
	}

gameloop:
	for {
		fmt.Println()

		if _, ok := Rooms[g.CurrentRoom]; !ok {
			log.Fatalf("transitioned to undefined Room #%03d", g.CurrentRoom)
		}

		// Setup AllCommands (needed for tab-completion)
		AllCommands = Rooms[g.CurrentRoom].Actions

		// Print the current Room's description
		if displayDescription {
			linenoise.Clear()
			fmt.Print(Rooms[g.CurrentRoom].Description(g))
		}
		displayDescription = false

		// Get the user's command
		line, err := linenoise.Line("> ")
		line = strings.ToLower(line)

		// Exit the game if the user pressed CTRL+C or CTRL+D
		if err == linenoise.KillSignalError {
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		linenoise.AddHistory(line)

		// Transform the input to catch slightly different ways of phrasing a command
		tline := transformCommand(line)

		if tline == "" {
			continue
		}

		if DEBUG && tline != line {
			log.Printf("Transformed input: %q\n\n", tline)
		}

		completion := CommandCompletion(tline)
		command := ""

		if len(completion) == 0 {
			fmt.Printf("\nSorry, I did not understand the command %q\n", line)
			continue
		} else if len(completion) == 1 {
			if completion[0] != tline {
				fmt.Println(">", completion[0])
			}
			command = completion[0]
		} else {
			fmt.Printf("\nThe command %q was ambiguous. Which of the following did you mean?\n\n", line)
			for _, c := range completion {
				fmt.Printf("* %s\n", c)
			}
			continue
		}

		switch command {
		case "exit", "quit":
			return

		case "help":
			fmt.Println(strings.TrimSpace(`
Blah blah blah, Tab completion, blah blah.

You can type in 'inventory' to examine your inventory.
			`))

		case "inventory":
			fmt.Println("\nInventory is not yet implemented")

		case "flags":
			if !DEBUG {
				fmt.Println("\nThis command is only available in DEBUG mode")
				continue
			}

			flags := make([]string, 0, len(g.flags))
			for k := range g.flags {
				flags = append(flags, k)
			}
			sort.Strings(flags)

			fmt.Println()
			for _, flag := range flags {
				fmt.Printf("* %s\n", flag)
			}

		case "commands":
			if !DEBUG {
				fmt.Println("\nThis command is only available in DEBUG mode")
				continue
			}

			var cmds []string

			for _, action := range AllCommands {
				cmds = append(cmds, action.Command[0])
			}
			sort.Strings(cmds)

			for _, cmd := range cmds {
				fmt.Printf("* %s\n", cmd)
			}

		case "xyzzy":
			fmt.Println("\nNothing happens.")

		case "die":
			fmt.Println("\nYou give up on life and wither away.")
			return

		default:
			for _, action := range Rooms[g.CurrentRoom].Actions {
				for _, cmd := range action.Command {
					if cmd == command {
						fmt.Println()
						displayDescription = Rooms[g.CurrentRoom].ExecuteAction(command, g)
						continue gameloop
					}
				}
			}
		}

	}
}
