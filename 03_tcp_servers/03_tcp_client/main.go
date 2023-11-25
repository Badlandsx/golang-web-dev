package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Juste write to tcp server who I am
	clientName := "client1"
	io.WriteString(conn, fmt.Sprintf("Hello from client %s", clientName))
	if err != nil {
		log.Fatalln(err)
	}
}
