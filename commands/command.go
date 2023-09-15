package commands

import (
	"serhhatsari/gores/constants"
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
	case constants.PingCmd:
		return handlePingCmd(command)
	case constants.SetCmd:
		return handleSetCmd(command)
	case constants.MSetCmd:
		return handleMsetCmd(command)
	case constants.GetCmd:
		return handleGetCmd(command)
	case constants.MGetCmd:
		return handleMgetCmd(command)
	case constants.CommandCmd:
		return handleCommandListCmd()
	case constants.DelCmd:
		return handleDelCmd(command)
	case constants.IncrCmd:
		return handleIncrCmd(command)
	case constants.DecrCmd:
		return handleDecrCmd(command)
	case constants.AppendCmd:
		return handleAppendCmd(command)
	default:
		return "-ERR Unknown command\r\n"
	}
}
