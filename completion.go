package main

import (
	"strings"
)

// CommandCompletion completes the given input
func CommandCompletion(input string) []string {
	if input == "" {
		return []string{"help", "commands", "inventory", "quit"}
	}

	var result []string

	tinput := transformCommand(strings.ToLower(input))

	for _, cmd := range AllCommands {
		if strings.HasPrefix(cmd, tinput) {
			result = append(result, cmd)
		}
	}

	return result
}
