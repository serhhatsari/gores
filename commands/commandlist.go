package commands

import (
	"fmt"
	"strings"
)

func handleCommandListCmd() string {
	// List of available commands
	commands := []string{
		"SET - Write data to the database",
		"GET - Read data from the database",
		"PING - Check if the server is alive",
		"DEL - Delete data from the database",
		"LPUSH - Insert all the specified values at the head of the list stored at key",
		"LPUSHX - Insert all the specified values at the head of the list stored at key, only if key already exists and holds a list",
		"LLEN - Returns the length of the list stored at key",
		"LPOP - Removes and returns the first element of the list stored at key",
		"LRANGE - Returns the specified elements of the list stored at key",
		"COMMAND - Returns details about all Redis commands",
		"STRLEN - Returns the length of the string value stored at key",
		"MSETNX - Sets the given keys to their respective values",
		"INCR - Increments the number stored at key by one",
		"INCRBY - Increments the number stored at key by increment",
		"INCRBYFLOAT - Increment the string representing a floating point number stored at key by the specified increment",
		"DECR - Decrements the number stored at key by one",
		"DECRBY - Decrements the number stored at key by decrement",
		"APPEND - Append a value to a key",
		"SETRANGE - Overwrites part of the string stored at key, starting at the specified offset",
		"GETRANGE - Returns the substring of the string value stored at key, determined by the offsets start and end",
		"GETDEL - Get the value of key and delete the key",
		"MGET - Returns the values of all specified keys",
		"GETEX - Get the value of key and set the key to a new value",
		"MSET - Sets the given keys to their respective values",
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
