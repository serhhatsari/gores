package commands

import "fmt"

func handleAppendCmd(command *Command) string {
	if len(command.Args) != 2 {
		return "-ERR APPEND requires 2 arguments\r\n"
	}

	key := command.Args[0]
	value := command.Args[1]

	currVal, ok := get(key)
	if !ok {
		set(key, value)
		return fmt.Sprintf(":%d\r\n", len(value))
	}

	set(key, currVal+value)
	return fmt.Sprintf(":%d\r\n", len(currVal+value))
}
