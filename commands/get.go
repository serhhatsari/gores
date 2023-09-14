package commands

import "fmt"

func get(key string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	value, ok := dataStore[key]
	return value, ok
}

func handleGetCmd(command *Command) string {
	if command.ArgsNum != 1 {
		return "-ERR Wrong number of arguments for GET\r\n"
	}

	key := command.Args[0]
	value, ok := get(key)
	if !ok {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}
