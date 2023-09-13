package commands

import (
	"sync"
)

type Command struct {
	Name    string
	ArgsNum int
	Args    []string
}

var (
	dataStore = make(map[string]string)
	mutex     = &sync.Mutex{}
)

func HandleCommand(command *Command) string {
	switch command.Name {
	case PingCmd:
		return handlePingCmd(command)
	case SetCmd:
		return handleSetCmd(command)
	case GetCmd:
		return handleGetCmd(command)
	case CommandCmd:
		return handleCommandListCmd()
	case DelCmd:
		return handleDelCmd(command)
	default:
		return "-ERR Unknown command\r\n"
	}
}
