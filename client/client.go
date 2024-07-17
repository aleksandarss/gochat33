package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	sock, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed connecting to the server; reason:", err)
	}

	defer sock.Close()
	fmt.Println("Connected to the server:", sock.RemoteAddr())

	var menuOption string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What do you want to do:\n1.register\n2.view list of available users [WIP]\n3.Start a chat with a user [WIP]\n")
	if scanner.Scan() {
		menuOption = scanner.Text()
	}

	switch menuOption {
	case "1":
		register(sock)
	case "2":
		fmt.Println("[INFO] This functionality is still in development")
	case "3":
		fmt.Println("[INFO] This functionality is still in development")
	default:
		fmt.Println("[ERROR] Unknown option:", menuOption)
	}
}

func register(sock net.Conn) {
	var username string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your desired username:")
	if scanner.Scan() {
		username = scanner.Text()
	}

	request := fmt.Sprintf("req:register;\nusername:%s", username)

	_, err := sock.Write([]byte(request))

	if err == nil {
		fmt.Printf("[INFO] Successfully registered with username: %s\n", username)
	}
}
