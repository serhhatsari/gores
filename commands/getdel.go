package commands

import (
	"fmt"
)

func handleGetDelCmd(command *Command) string {
	if command.ArgsNum != 1 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	value, ok := get(key)
	if !ok {
		return "$-1\r\n"
	}

	del(key)

	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}
