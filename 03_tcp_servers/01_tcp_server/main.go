package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen on port 8080
	// Use telnet localhost 8080 to test
	tcpServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpServer.Close()

	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
