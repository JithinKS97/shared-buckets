package main

import (
	"bufio"
	"fmt"
	"os"

	nkn "github.com/nknorg/nkn-sdk-go"
)

func createClient() {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		panic("Unable to create account")
	}

	client, err := nkn.NewClient(account, "", nil)

	if err != nil {
		panic("Unable to create a client")
	}

	fmt.Println("Created a client with address")
	fmt.Println(client.Address())

	<-client.OnConnect.C
	fmt.Println("Connected to the client")
	msg := make(chan nkn.Message)

	go func() {
		for {
			msg := <-client.OnMessage.C
			println(string(msg.Data))
		}
	}()

	<-msg
}

func connectToClient() {
	fmt.Print("Enter the address of the client to connect to\n")
	var address string
	fmt.Scanln(&address)

	account, err := nkn.NewAccount(nil)
	if err != nil {
		panic("Unable to create account")
	}

	client, err := nkn.NewClient(account, "", nil)

	if err != nil {
		panic("Unable to create a client")
	}

	for {
		fmt.Print("Enter the message to send\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := scanner.Text()
		client.Send(nkn.NewStringArray(address), []byte(message), nil)
	}
}
