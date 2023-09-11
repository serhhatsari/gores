package commands

func handlePingCmd(command *Command) string {
	if command.ArgsNum == 2 {
		return "+" + command.Args[0] + "\r\n"
	}
	return "+PONG\r\n"
}
