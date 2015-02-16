package main

import (
	"strings"
)

// CommandStrings is the list of top-level commands
var CommandStrings = []string{"help", "commands", "inventory", "exit", "quit"}

// CommandCompletion completes the given input
func CommandCompletion(input string) []string {
	if input == "" {
		return CommandStrings
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
