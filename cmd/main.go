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
		if len(args) < 2 {
			fmt.Println("Usage: netcli ping <host>")
			os.Exit(1)
		}
		runPing(args[1])
	case "lookup":
		if len(os.Args) < 3 {
			fmt.Println("Usage: netcli lookup <hostname>")
			os.Exit(1)
		}
		runLookup(args[1])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

}

// func runSystemPing(host string) error {
// 	// Windows uses "ping -n 4", Linux/Mac uses "ping -c 4"
// 	cmd := exec.Command("ping", "-n", "4", host) // change -n to -c if on Linux/Mac
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }
