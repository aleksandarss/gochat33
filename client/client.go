package main

import (
	"fmt"
	"net"
)

func main() {
	sock, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed connecting to the server; reason:", err)
	}

	defer sock.Close()
	fmt.Println("Connected to the server:", sock.RemoteAddr())
	_, err = sock.Write([]byte("Hello there server! How is it going?"))
}
