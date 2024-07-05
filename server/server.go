package main

import (
	"fmt"
	"net"
)

const PORT string = "8080"
const HOST string = "localhost"
const CONN_TYPE string = "tcp"

func main() {
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
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	messageLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed reading message; reason:", err)
		panic(err)
	}
	fmt.Println("Received message:", string(buffer[:messageLen]))
}
