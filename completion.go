package main

import (
	"strings"
)

// CommandStrings is the list of top-level commands
var CommandStrings = []string{"help", "inventory", "exit", "quit", "flags", "commands"}

func allCommandsHas(command string) bool {
	for _, action := range AllCommands {
		for _, cmd := range action.Command {
			if cmd == command {
				return true
			}
		}
	}
	return false
}

// CommandCompletion completes the given input
func CommandCompletion(input string) []string {
	if input == "" {
		return CommandStrings
	}

	if input == "xyzzy" {
		return []string{"xyzzy"}
	}

	if input == "die" || input == "kill yourself" {
		return []string{"die"}
	}

	if len(input) == 1 {
		switch input {
		case "n":
			if allCommandsHas("go north") {
				return []string{"go north"}
			}

		case "e":
			if allCommandsHas("go east") {
				return []string{"go east"}
			}

		case "s":
			if allCommandsHas("go south") {
				return []string{"go south"}
			}

		case "w":
			if allCommandsHas("go west") {
				return []string{"go west"}
			}
		}
	}

	var result []string

	tinput := transformCommand(strings.ToLower(input))

	for _, tlc := range CommandStrings {
		if strings.HasPrefix(tlc, tinput) {
			result = append(result, tlc)
		}
	}

	for _, action := range AllCommands {
		for _, cmd := range action.Command {
			if strings.Contains(cmd, tinput) {
				result = append(result, action.Command[0])
				break
			}
		}
	}

	return result
}
