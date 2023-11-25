package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Use telnet localhost 8080 to connect to tcp server
	// You can use http://localhost:8080 to see that HTTP is superset of TCP
	tcpServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpServer.Close()

	// Accept all connections
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Put timeout deadline on connection
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	// net.Conn is an io.Reader because of implementing the interface io.Reader (Read method)
	scanner := bufio.NewScanner(conn)
	// Allow to scan all client "request"
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "Hello" {
			io.WriteString(conn, "Hello from tcp server")
		} else {
			io.WriteString(conn, "You should say 'Hello' ......")
		}
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
