package main

import (
	"fmt"
	"log"
	"net"
)


const (
	// ServerAddress - address of the server
	ServerAddress = ":8090"
)

// receiver go routine
func startReceiver(conn net.Conn) {
	/* Read and print client message */
	b := make([]byte, 1024)
	for {
		n, err := conn.Read(b)
		if err != nil {
			log.Printf("Closing receiver with error: %v \n",err)
			return
		}
		fmt.Print(string(b[:n]))
	}
}

func main() {
	/* open a listener */
	listener, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatalf("Error starting servcer: %v", err)
	}
	fmt.Println("Server running on " + listener.Addr().String())
	for {
		/* Wait for a new connection */
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepitng connection: %v", err)
		}
		/* Create a new go routine for processing the new client connection */
		go startReceiver(conn)
	}
}
