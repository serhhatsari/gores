package commands

func set(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	dataStore[key] = value
}

func handleSetCmd(command *Command) string {
	if command.ArgsNum != 2 {
		return "-ERR Wrong number of arguments for SET\r\n"
	}

	key := command.Args[0]
	value := command.Args[1]

	set(key, value)

	return "+OK\r\n"
}
