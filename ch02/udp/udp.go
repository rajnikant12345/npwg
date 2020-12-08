package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len( os.Args) < 2 {
		log.Fatalf("Please enter the sender address")
	}
	serverAddr := os.Args[1]
	clientAddr := os.Args[2]

	sa,err := net.ResolveUDPAddr("udp",serverAddr)
	if err != nil {
		log.Fatalf(" Listen closed with error: %v", err)
	}

	l, err := net.ListenUDP("udp",sa)

	if err != nil {
		log.Fatalf(" Listen closed with error: %v", err)
	}

	ca, err := net.ResolveUDPAddr("udp",clientAddr)
	if err != nil {
		log.Fatalf(" Listen closed with error: %v", err)
	}

	fmt.Println("Enter text to send message.................. ")

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			if strings.HasPrefix(text, "exit") {
				fmt.Println("Exiting client loop ...")
				break
			}
			_,err := l.WriteToUDP( []byte(text), ca )
			if err != nil {
				log.Printf(" Listen closed with error: %v\n", err)
				break
			}

		}
	}()

	message := make([]byte,1024)

	for {
		n,add,err := l.ReadFrom(message)
		if err != nil {
			log.Printf(" Receiver closed with error: %v\n", err)
			break
		}
		fmt.Println("======================================================")
		fmt.Printf("Bytes Received: %d \n",n)
		fmt.Printf("Address of sender: %s \n",add.String())
		fmt.Printf("Message: %s",string(message[:n]))
		fmt.Println("======================================================")

	}
}
