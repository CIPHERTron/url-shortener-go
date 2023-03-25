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

		go handleConnection(conn)
	}
}

func handleConnection(src net.Conn) {
	destination, err := net.Dial("tcp", "https://www.google.com:80")

	if err != nil {
		log.Fatalln("Unable to connect to target server")
	}

	defer destination.Close()
}
