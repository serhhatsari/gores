package commands

func handleMsetCmd(command *Command) string {
	if command.ArgsNum == 0 || command.ArgsNum%2 != 0 {
		return "-ERR Wrong number of arguments for MSET\r\n"
	}

	for i := 0; i < command.ArgsNum; i += 2 {
		key := command.Args[i]
		value := command.Args[i+1]
		set(key, value)
	}

	return "+OK\r\n"
}
