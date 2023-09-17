package commands

func handleMSetNxCmd(command *Command) string {
	if command.ArgsNum%2 != 0 {
		return "-ERR wrong number of arguments for MSETNX\r\n"
	}

	for i := 0; i < command.ArgsNum; i += 2 {
		key := command.Args[i]
		value := command.Args[i+1]
		if _, ok := get(key); ok {
			return ":0\r\n"
		}
		set(key, value)
	}

	return ":1\r\n"
}
