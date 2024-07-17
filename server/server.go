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
	src        string
	dst        string
	msgcontent string
	req        string
	username   string
}

const PORT string = "8080"
const HOST string = "localhost"
const CONN_TYPE string = "tcp"

/*
PROTOCOL NOTES:
1. Server starts
2. Client 1 connects and publishes its username -> "req:register;\nusername:<some_username>"
3. Client 2 connects and also publishes its username -> "req:register;\nusername:<second_username>"
4. Either client sends a message -> "req:send;\nsrc:<username>;\ndst:<username>;\nmsg:<message to the other client>"
*/

func main() {
	clients := make([]Client, 0)

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

	message := getRequestParams(&buffer)

	fmt.Println("Request params: ", message)

	switch message.req {
	case "register":
		registerNewClient(message.username, conn.RemoteAddr().String(), clients)
	default:
		fmt.Printf("[INFO] Functionality for request = [%s] not implemented yet\n", message.req)
	}

	fmt.Println("Registered clients:", clients)
}

func getRequestParams(request *[]byte) (message Message) {
	messageStr := string(*request)

	params := strings.Split(messageStr, ";\n")

	for _, param := range params {
		paramParts := strings.Split(param, ":")

		switch key := paramParts[0]; key {
		case "req":
			message.req = paramParts[1]
		case "username":
			message.username = paramParts[1]
		case "src":
			message.src = paramParts[1]
		case "dst":
			message.dst = paramParts[1]
		case "msgcontent":
			message.msgcontent = paramParts[1]
		default:
			fmt.Printf("[ERROR] (getRequstType): param = [%s] not accepted.\n", key)
		}
	}

	return message
}

func registerNewClient(username string, address string, clients *[]Client) {
	*clients = append(*clients, Client{
		username: username,
		host:     strings.Split(address, ":")[0],
		port:     strings.Split(address, ":")[1],
	})
}

func getClientByUsername(username string, clients *[]Client) (client Client) {
	for _, client := range *clients {
		if client.username == username {
			return client
		}
	}
	return client
}

func sendMessage(srcClient, dstClient *Client, message string) {

}
