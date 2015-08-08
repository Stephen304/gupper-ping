package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

func main() {
	defer os.Exit(1)

	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("{\"Latency\": %v}\n", rtt.Seconds()*1000)
		os.Exit(0)
	}
	p.Run()
}
