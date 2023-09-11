package main

import "serhhatsari/gores/commands"

func executeCommand(command *commands.Command) string {
	switch command.Name {
	case commands.PingCmd:
		return commands.ExecutePingCommand(command)
	case commands.SetCmd:
		return commands.ExecuteSetCommand(command)
	case commands.GetCmd:
		return commands.ExecuteGetCommand(command)
	case commands.CommandCmd:
		return commands.ExecuteCommandList()
	default:
		return "-ERR Unknown command\r\n"
	}
}
