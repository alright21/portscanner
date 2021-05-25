package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

var domain string = "scanme.nmap.org"

func worker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", domain, port)

		connection, scanError := net.DialTimeout("tcp", address, 500*time.Millisecond)

		if scanError != nil {
			// the port is closed or filtered
			results <- 0
			continue
		}

		connection.Close()
		results <- port

	}
}
func main() {

	ports := make(chan int, 100)

	results := make(chan int)

	openPorts := make([]int, 0)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for port := 1; port <= 1024; port++ {
			ports <- port
		}
	}()

	// results will be filled 1024 times, so I can use this fixed loop
	for i := 0; i < 1024; i++ {
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("Port %d: open\n", port)
	}

	fmt.Printf("Scan result (1-1024):\n\n%s has %d open ports\n", domain, len(openPorts))
}
