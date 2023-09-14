package commands

import "fmt"

func handlePingCmd(command *Command) string {
	if command.ArgsNum == 2 {
		return fmt.Sprintf("+%s\r\n", command.Args[0])
	}
	return "+PONG\r\n"
}
