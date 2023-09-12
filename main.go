package main

import (
	"log/slog"
	"net"

	"serhhatsari/gores/commands"
	"serhhatsari/gores/utils"
)

func main() {
	slog.Info("Redis server is starting...")

	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		slog.Error("Error listening:", err.Error())
		return
	}
	// Close the listener when the application closes
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			slog.Error("Error closing listener:", err)
			return
		}
	}(listener)

	slog.Info("Server is initialized")
	slog.Info("Ready to accept connections tcp")

	for {
		// Accept a connection from a client
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Error accepting: ", err.Error())
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
			slog.Error("Error closing connection:", err)
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
			slog.Error("Error reading:", err.Error())
			return
		}

		// Convert the data to a string
		request := string(buf[:n])

		// Convert the request to a command
		command := utils.ConvertToCommand(request)

		// Parse and execute Redis command
		response := commands.HandleCommand(command)

		// Send the response back to the client
		_, err = conn.Write([]byte(response))
		if err != nil {
			slog.Error("Error writing:", err.Error())
			return
		}
	}
}
