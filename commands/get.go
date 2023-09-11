package commands

import "fmt"

func Get(key string) (string, bool) {
	mutex.Lock()
	value, ok := dataStore[key]
	mutex.Unlock()
	return value, ok
}

func handleGetCmd(command *Command) string {
	if command.ArgsNum != 1 {
		return "-ERR Wrong number of arguments for GET\r\n"
	}

	key := command.Args[0]
	value, ok := Get(key)
	if !ok {
		return "$-1\r\n"
	}
	return "$" + fmt.Sprint(len(value)) + "\r\n" + value + "\r\n"
}
