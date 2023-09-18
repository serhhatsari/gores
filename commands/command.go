package commands

import (
	"serhhatsari/gores/constants"
	"serhhatsari/gores/pkg"
	"sync"
)

type Command struct {
	Name    string
	ArgsNum int
	Args    []string
}

var (
	dataStore = make(map[string]string)
	listStore = make(map[string]*pkg.LinkedList)
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
	case constants.GetDelCmd:
		return handleGetDelCmd(command)
	case constants.DecrByCmd:
		return handleDecrByCmd(command)
	case constants.IncrByCmd:
		return handleIncrByCmd(command)
	case constants.GetRangeCmd:
		return handleGetRangeCmd(command)
	case constants.StrlenCmd:
		return handleStrlenCmd(command)
	case constants.SetRangeCmd:
		return handleSetRangeCmd(command)
	case constants.GetExCmd:
		return handleGetExCmd(command)
	case constants.IncrByFloatCmd:
		return handleIncrByFloatCmd(command)
	case constants.MSetNx:
		return handleMSetNxCmd(command)
	case constants.LPushCmd:
		return handleLPushCmd(command)
	case constants.LLenCmd:
		return handleLLenCmd(command)
	case constants.LPopCmd:
		return handleLPopCmd(command)
	default:
		return "-ERR Unknown command\r\n"
	}
}
