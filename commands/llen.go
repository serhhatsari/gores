package commands

import "fmt"

func handleLLenCmd(command *Command) string {
	if command.ArgsNum != 1 {
		return "-ERR wrong number of arguments for command\r\n"
	}

	key := command.Args[0]

	_, ok := listStore[key]
	if !ok {
		return ":0\r\n"
	}

	return fmt.Sprintf(":%d\r\n", llen(key))
}
