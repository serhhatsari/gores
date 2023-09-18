package commands

import (
	"fmt"
	"serhhatsari/gores/pkg"
)

func rpush(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := listStore[key]; !ok {
		listStore[key] = pkg.NewLinkedList()
	}
	listStore[key].PushBack(value)
}

func handleRPushCmd(command *Command) string {
	if command.ArgsNum < 2 {
		return "-ERR wrong number of arguments for RPush\r\n"
	}

	key := command.Args[0]
	for i := 1; i < command.ArgsNum; i++ {
		value := command.Args[i]
		rpush(key, value)
	}

	return fmt.Sprintf(":%d\r\n", llen(key))
}
