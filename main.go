package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	instances "public-cloud-api/cmd/instances"
)

type commandFunc func(args []string) error

var commands = map[string]commandFunc{
	"instances": handleInstances,
	// Future commands can be added here, e.g., "create": handleCreate,
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	// Ensure at least one command is provided
	if len(args) < 1 {
		fmt.Println("Error: A command is required.\nAvailable commands:", availableCommands())
		return
	}

	// Look up and execute the appropriate command
	cmd, found := commands[args[0]]
	if !found {
		fmt.Printf("Error: Unknown command '%s'.\nAvailable commands: %s\n", args[0], availableCommands())
		return
	}

	// Execute the command and handle any errors
	if err := cmd(args[1:]); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func handleInstances(args []string) error {
	if len(args) < 1 {
		return errors.New("a subcommand is required for 'instances'.\nAvailable subcommands: list")
	}

	// Subcommand handler for "instances"
	switch args[0] {
	case "list":
		return instances.ListInstances(args[1:]) // Delegate to the ListInstances function
	default:
		return fmt.Errorf("unknown subcommand '%s' for 'instances'.\nAvailable subcommands: list", args[0])
	}
}

// Utility function to list all available commands
func availableCommands() string {
	var cmds []string
	for cmd := range commands {
		cmds = append(cmds, cmd)
	}
	return strings.Join(cmds, ", ")
}
