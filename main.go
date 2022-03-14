package main

func main() {
	input := getCommandlineInput()
	switch input {
	case "connect":
		connectToClient()
	case "create":
		createClient()
	}
}
