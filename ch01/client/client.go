package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	//ServerAddress - address of the server to connect
	ServerAddress = "127.0.0.1:8090"
)


func main() {
	/******** Section for dialing to a server ***************/

	client, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		log.Fatalf("Error while opening connection: %v",err)
	}
	defer func() {
		err = client.Close()
		log.Fatalf("Error while closing connection: %v",err)
	}()

	/***********************************************************/

	/* loop until exit condition is hit */
	for {
		/****************** Read message Begin ********************/

		fmt.Print("Type Client Message")
		fmt.Print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		/******************* Read Message End *********************/

		/***************** Check exit condition **********************/
		if strings.HasPrefix(text, "exit") {
			fmt.Println("closing connection...")
			break
		}
		/************************************************************/

		/********************** write message to server ****************/
		_, err = client.Write([]byte(text))
		if err != nil {
			log.Fatalf("Error while writing over connection: %v",err)
		}
		/*************************************************************/
	}
}
