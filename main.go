package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalln("Failed to establish connection in port 5000")
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
	}
}
