package main

import (
	"fmt"
	"os"

	"github.com/toba-madamori/netcli/internal/netutils"
)

func runLookup(host string) {
	cname, ips, err := netutils.LookupHost(host)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Canonical name:", cname)
	fmt.Println("IP addresses:")
	for _, ip := range ips {
		fmt.Println(" -", ip)
	}
}
