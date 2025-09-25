package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toba-madamori/netcli/internal/netutils"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup [hostname]",
	Short: "Lookup DNS canonical name and IP addresses for a hostname",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		cname, ips, err := netutils.LookupHost(host)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Canonical name:", cname)
		fmt.Println("IP addresses:")
		for _, ip := range ips {
			fmt.Println(" -", ip)
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
}
