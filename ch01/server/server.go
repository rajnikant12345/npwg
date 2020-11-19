package main

import (
	"fmt"
	"net"
)

const (
	ServerAddress = ":8090"
)

func startReceiver(conn net.Conn) {
	b := make([]byte,1024)
	for {
		n,err := conn.Read(b)
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println(string(b[:n]))
	}
}


func main() {
	listener,err := net.Listen("tcp",ServerAddress)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		startReceiver(conn)
	}
}


