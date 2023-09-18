package commands

func handleRPushXCmd(command *Command) string {
	if command.ArgsNum < 2 {
		return "-ERR wrong number of arguments for RPush\r\n"
	}

	key := command.Args[0]
	if !exists(key) {
		return ":0\r\n"
	}

	return handleRPushCmd(command)
}
