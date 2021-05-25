package main

import (
	"fmt"
	"log"
	"net"
)

func main(){

	_, scanError := net.Dial("tcp", "scanme.nmap.org:80")

	if scanError != nil {
		log.Fatal("Scan Error: ", scanError)
	}

	fmt.Println("Connection successful")
}
