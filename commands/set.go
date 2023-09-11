package commands

func Set(key string, value string) {
	mutex.Lock()
	dataStore[key] = value
	mutex.Unlock()
}

func ExecuteSetCommand(command *Command) string {
	if command.ArgsNum != 2 {
		return "-ERR Wrong number of arguments for SET\r\n"
	}

	key := command.Args[0]
	value := command.Args[1]

	Set(key, value)

	return "+OK\r\n"
}
