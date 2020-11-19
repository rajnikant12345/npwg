package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
)

const (
	ServerAddress = "localhost:8090"
)

func getClientID() string {
	b := make([]byte,8)
	_,err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func main() {
	client, err := net.Dial("tcp",ServerAddress)
	if err != nil {
		panic(err)
	}

	defer client.Close()
	clientID := getClientID()
	for {

		fmt.Print("Type Client Message: ")
		in,_ := ioutil.ReadAll(os.Stdin)
		if string(in) == "exit" {
			break
		}
		_, err = client.Write([]byte("ClientID: "+ clientID + " Message: " + string(in) ))
		if err != nil {
			panic(err)
		}
	}
}
