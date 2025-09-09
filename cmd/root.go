package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is always the program name(netcli)
	// os.Args[1:] are the actual parameters

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: netcli <command> [arguments]")
		os.Exit(1)
	}

	// First argument is the command (e.g. ping, lookup)

	command := args[0]

	switch command {
	case "ping":
		fmt.Println("You called ping!")
	case "lookup":
		fmt.Println("You called lookup!")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

}
