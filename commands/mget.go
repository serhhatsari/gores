package commands

import "fmt"

func handleMgetCmd(command *Command) string {
	if command.ArgsNum == 0 {
		return "-ERR Wrong number of arguments for MGET\r\n"
	}

	var values string

	for i := 0; i < command.ArgsNum; i++ {
		key := command.Args[i]
		value, ok := get(key)
		if !ok {
			values += "$-1\r\n"
		} else {
			values += fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
		}
	}

	return fmt.Sprintf("*%d\r\n%s", command.ArgsNum, values)
}
