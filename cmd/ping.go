package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/toba-madamori/netcli/internal/netutils"
)

var pingCmd = &cobra.Command{
	Use:   "ping [host]",
	Short: "Send ICMP echo requests to host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		method, _ := cmd.Flags().GetString("method") // "go" or "system"

		switch method {
		case "go":
			// Pure Go ICMP (may require admin privileges)
			if err := netutils.GoPing(host); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		case "system":
			// Use OS ping binary
			if err := runSystemPing(host); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		default:
			fmt.Fprintln(os.Stderr, "Unknown method:", method)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
	// flag: --method, -m (default "go")
	pingCmd.Flags().StringP("method", "m", "go", "ping method: 'go' (raw ICMP) or 'system' (call OS ping)")
}

// runSystemPing calls the OS ping command (cross-platform)
func runSystemPing(host string) error {
	var args []string
	if runtime.GOOS == "windows" {
		args = []string{"-n", "4", host}
	} else {
		args = []string{"-c", "4", host}
	}
	cmd := exec.Command("ping", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
