package main

import (
	"fmt"
	"strings"
)

const (
	PingCmd    = "PING"
	SetCmd     = "SET"
	GetCmd     = "GET"
	CommandCmd = "COMMAND"
)

func Set(key string, value string) {
	mutex.Lock()
	dataStore[key] = value
	mutex.Unlock()
}

func Get(key string) (string, bool) {
	mutex.Lock()
	value, ok := dataStore[key]
	mutex.Unlock()
	return value, ok
}

func executeRedisCommand(command *Command) string {

	switch command.Name {

	case PingCmd:
		if command.ArgsNum == '2' {
			return "+" + command.Args[0] + "\r\n"
		}
		return "+PONG\r\n"

	case SetCmd:
		if command.ArgsNum != 2 {
			return "-ERR Wrong number of arguments for SET\r\n"
		}

		key := command.Args[0]
		value := command.Args[1]

		Set(key, value)

		return "+OK\r\n"
	case GetCmd:
		if command.ArgsNum != 1 {
			return "-ERR Wrong number of arguments for GET\r\n"
		}
		key := command.Args[0]
		value, ok := Get(key)
		if !ok {
			return "$-1\r\n"
		}
		return "$" + fmt.Sprint(len(value)) + "\r\n" + value + "\r\n"
	case CommandCmd:
		return formatCommandList()
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
