package commands

import (
	"fmt"
	"serhhatsari/gores/pkg"
)

func lLen(key string) int {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := listStore[key]; !ok {
		return 0
	}
	return listStore[key].Size()
}

func lpush(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := listStore[key]; !ok {
		listStore[key] = pkg.NewLinkedList()
	}
	listStore[key].PushFront(value)
}

func handleLPushCmd(command *Command) string {
	if command.ArgsNum < 2 {
		return "-ERR wrong number of arguments for LPush\r\n"
	}

	key := command.Args[0]
	for i := 1; i < command.ArgsNum; i++ {
		value := command.Args[i]
		lpush(key, value)
	}

	return fmt.Sprintf(":%d\r\n", lLen(key))
}
