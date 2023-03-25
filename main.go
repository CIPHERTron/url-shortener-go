package main

import (
	"io"
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

	go func() {
		if _, err := io.Copy(destination, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, destination); err != nil {
		log.Fatalln(err)
	}
}
