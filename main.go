package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	dataStore = make(map[string]string)
	mutex     = &sync.Mutex{}
)

type Command struct {
	Name    string
	ArgsNum int
	Args    []string
}

func main() {
	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	// Close the listener when the application closes
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing listener:", err)
			return
		}
	}(listener)

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
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
			return
		}
	}(conn)

	for {
		// Create a buffer to hold incoming data
		buf := make([]byte, 1024)

		// Read data from the connection
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			fmt.Println("Error reading:", err)
			return
		}

		// Convert the data to a string
		request := string(buf[:n])

		// Convert the request to a command
		command := convertToCommand(request)

		// Parse and execute Redis command
		response := executeRedisCommand(command)

		// Send the response back to the client
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func convertToCommand(request string) *Command {
	parts := strings.Fields(request)

	name := strings.ToUpper(parts[2])

	argsNum := int(parts[0][1]-'0') - 1

	var args []string
	for i := 4; i < len(parts); i += 2 {
		args = append(args, parts[i])
	}

	return &Command{
		Name:    name,
		ArgsNum: argsNum,
		Args:    args,
	}
}
