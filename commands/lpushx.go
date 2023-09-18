package commands

func exists(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := listStore[key]
	return ok
}

func handleLPushXCmd(command *Command) string {
	if len(command.Args) == 0 {
		return "ERR wrong number of arguments for '" + command.Name + "' command"
	}

	key := command.Args[0]
	if !exists(key) {
		return ":0\r\n"
	}

	return handleLPushCmd(command)
}
