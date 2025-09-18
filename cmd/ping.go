package main

import (
	"fmt"

	"github.com/toba-madamori/netcli/internal/netutils"
)

func runPing(host string) {
	err := netutils.GoPing(host)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
