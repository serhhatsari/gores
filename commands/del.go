package commands

import "fmt"

func del(key string) bool {
	mutex.Lock()
	_, ok := dataStore[key]
	if ok {
		delete(dataStore, key)
	}
	mutex.Unlock()
	return ok
}

func handleDelCmd(command *Command) string {
	if command.ArgsNum < 1 {
		return "-ERR wrong number of arguments for 'del' command\r\n"
	}

	numOfDeletedKeys := 0

	for _, key := range command.Args {
		isDeleted := del(key)
		if isDeleted {
			numOfDeletedKeys++
		}
	}

	if numOfDeletedKeys == 1 {
		return "+OK\r\n"
	}

	return ":" + fmt.Sprint(numOfDeletedKeys) + "\r\n"
}
