package commands

import "strconv"

func handleIncrByCmd(command *Command) string {
	if command.ArgsNum != 2 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]
	incr, err := strconv.Atoi(command.Args[1])
	if err != nil {
		return "-ERR Value is not an integer\r\n"
	}

	return increment(key, incr)
}
