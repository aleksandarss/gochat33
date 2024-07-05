package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	username string
	host     string
	port     string
}

type Message struct {
	src     string
	dst     string
	content string
}

const PORT string = "8080"
const HOST string = "localhost"
const CONN_TYPE string = "tcp"

/*
PROTOCOL NOTES:
1. Server starts
2. Client 1 connects and publishes its username -> "username:<some_username>"
3. Client 2 connects and also publishes its username -> "username:<second_username>"
4. Either client sends a message -> "src:<username>;dst:<username>;msg:<message to the other client>"
*/

func main() {
	clients := make([]Client, 8)

	sock, err := net.Listen(CONN_TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Failed creating socket; reason:", err)
		return
	}

	defer sock.Close()
	fmt.Printf("Server listening on port = [%s]\n", PORT)

	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Failed accepting connection; reason:", err)
			continue
		}
		handleConnection(conn, &clients)
	}
}

func handleConnection(conn net.Conn, clients *[]Client) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	messageLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed reading message; reason:", err)
		panic(err)
	}
	fmt.Println("Received message:", string(buffer[:messageLen]))
	if strings.Index(string(buffer[:messageLen]), "username:") != -1 {
		fmt.Printf("Client at addr: [%s] wants to register to the chat ...", conn.RemoteAddr())
		registerNewClient(&conn, clients)
	}
}

func getMessageParams(message *[]byte) {

}

func registerNewClient(conn *net.Conn, clients *[]Client) {

}
