package utils

import (
	"strings"

	"serhhatsari/gores/commands"
)

func ConvertToCommand(request string) *commands.Command {
	parts := strings.Fields(request)

	name := strings.ToUpper(parts[2])

	argsNum := int(parts[0][1]-'0') - 1

	var args []string
	for i := 4; i < len(parts); i += 2 {
		args = append(args, parts[i])
	}

	return &commands.Command{
		Name:    name,
		ArgsNum: argsNum,
		Args:    args,
	}
}