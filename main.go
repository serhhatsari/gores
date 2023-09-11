package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	dataStore = make(map[string]string)
	mutex     = &sync.Mutex{}
)

func main() {
	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	// Close the listener when the application closes
	defer listener.Close()

	for {
		// Accept a connection from a client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		// Create a buffer to hold incoming data
		buf := make([]byte, 1024)

		// Read data from the connection
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Convert the data to a string
		request := string(buf[:n])

		// Parse and execute Redis command
		response := executeRedisCommand(request)

		// Send the response back to the client
		conn.Write([]byte(response))
	}
}
