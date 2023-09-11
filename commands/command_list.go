package commands

import (
	"fmt"
	"strings"
)

func ExecuteCommandList() string {
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
