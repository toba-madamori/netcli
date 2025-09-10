package netutils

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func GoPing(host string) error {
	// Resolve hostname to IP
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return fmt.Errorf("failed to resolve host: %v", err)
	}

	// Create ICMP connection
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer conn.Close()

	// Build ICMP Echo request
	icmpMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("Hello-Ping"),
		},
	}

	msgBytes, err := icmpMessage.Marshal(nil)
	if err != nil {
		return err
	}

	// Send packet
	start := time.Now()
	_, err = conn.WriteTo(msgBytes, ipAddr)
	if err != nil {
		return err
	}

	// Receive reply
	reply := make([]byte, 1500)
	err = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		return err
	}

	n, _, err := conn.ReadFrom(reply)
	if err != nil {
		return fmt.Errorf("request timeout: %v", err)
	}

	// Parse response
	duration := time.Since(start)
	parsedMessage, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		return err
	}

	switch parsedMessage.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("Reply from %s: time=%v\n", ipAddr.String(), duration)
	default:
		fmt.Printf("Got %+v from %s\n", parsedMessage, ipAddr.String())
	}

	return nil

}
