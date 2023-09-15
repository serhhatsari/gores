package commands

import (
	"fmt"
	"strconv"
)

func increment(key string, decr int) string {
	mutex.Lock()
	defer mutex.Unlock()

	value, ok := dataStore[key]
	if !ok {
		dataStore[key] = "0"
		value = "0"
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return "-ERR Value is not an integer\r\n"
	}
	intVal += decr

	dataStore[key] = strconv.Itoa(intVal)

	return fmt.Sprintf(":%d\r\n", intVal)
}

func handleIncrCmd(command *Command) string {
	if len(command.Args) != 1 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	return increment(key, 1)
}
