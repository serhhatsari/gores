package commands

import "strconv"

func handleStrlenCmd(command *Command) string {
	if command.ArgsNum != 1 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	val, ok := get(key)
	if !ok {
		return ":0\r\n"
	}

	return ":" + strconv.Itoa(len(val)) + "\r\n"
}
