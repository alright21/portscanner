package main

import (
	"fmt"
	"net"
)

func main(){

	openPorts := 0

	domain := "scanme.nmap.org"

	for port :=1; port <= 1024; port++{

		address := fmt.Sprintf("%s:%d",domain, port)

		connection, scanError := net.Dial("tcp", address)

		if scanError != nil {
			// the port is closed
			fmt.Printf("Port %d: closed\n", port)
			continue
		}

		fmt.Printf("Port %d: open\n", port)
		openPorts++
		connection.Close()

	}

	fmt.Printf("Scan result (1-1024):\n\n%s has %d open ports\n", domain, openPorts)
}
