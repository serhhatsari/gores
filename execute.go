package main

import (
	"fmt"
	"strings"
)

func executeRedisCommand(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "-ERR Empty command\r\n"
	}

	cmd := strings.ToUpper(parts[2])

	switch cmd {
	case "COMMAND":
		return formatCommandList()
	case "SET":
		if parts[0][1] != '3' {
			return "-ERR Wrong number of arguments for SET\r\n"
		}

		key := parts[4]
		value := parts[6]
		fmt.Println("SET", key, value)

		mutex.Lock()
		dataStore[key] = value
		mutex.Unlock()
		return "+OK\r\n"
	case "GET":
		if parts[0][1] != '2' {
			return "-ERR Wrong number of arguments for GET\r\n"
		}
		key := parts[4]
		fmt.Println("GET", key)

		mutex.Lock()
		value, ok := dataStore[key]
		mutex.Unlock()

		if !ok {
			return "$-1\r\n"
		}
		return "$" + fmt.Sprint(len(value)) + "\r\n" + value + "\r\n"

	case "PING":
		if parts[0][1] == '2' {
			return "+" + parts[4] + "\r\n"
		}
		return "+PONG\r\n"
	default:
		return "-ERR Unknown command\r\n"
	}
}

func formatCommandList() string {
	// List of available commands
	commands := []string{
		"SET - Write data to the database",
		"GET - Read data from the database",
		"PING - Check if the server is alive",
	}

	// Build the response in Redis protocol format
	var response strings.Builder
	response.WriteString("*" + fmt.Sprint(len(commands)) + "\r\n")

	for _, cmdDesc := range commands {
		response.WriteString("$" + fmt.Sprint(len(cmdDesc)) + "\r\n")
		response.WriteString(cmdDesc + "\r\n")
	}

	return response.String()
}
