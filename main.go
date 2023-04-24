package main

import (
	"fmt"
	"net"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	ip := "localhost"
	startPort := 1
	endPort := 65535
	bar := progressbar.NewOptions(endPort-startPort,
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[cyan][Port Scan]"),
	)
	openPorts := []int{}

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, time.Millisecond*500)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
		bar.Add(1)
	}

	for _, port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}
}
