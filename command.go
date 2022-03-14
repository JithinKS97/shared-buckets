package main

import "fmt"

func getCommandlineInput() string {
	fmt.Print("1. Create\n2. Connect\n")
	fmt.Print("Enter an option\n")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		input = "create"
	case "2":
		input = "connect"
	}
	return input
}
