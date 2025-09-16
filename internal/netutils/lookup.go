package netutils

import (
	"fmt"
	"net"
)

// LookupHost resolves the IP addresses and CNAME of a given host
func LookupHost(host string) (string, []string, error) {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return "", nil, fmt.Errorf("CNAME lookup failed: %w", err)
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return "", nil, fmt.Errorf("IP lookup failed: %w", err)
	}

	var results []string
	for _, ip := range ips {
		results = append(results, ip.String())
	}

	return cname, results, nil
}
