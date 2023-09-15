package utils

import (
	"fmt"
	"strconv"
	"strings"

	"serhhatsari/gores/commands"
)

func ConvertToCommand(request string) *commands.Command {

	parts := strings.Fields(request)

	name := strings.ToUpper(parts[2])

	argsNum, err := strconv.Atoi(parts[0][1:])
	if err != nil {
		fmt.Println("Error converting args num to int")
		return nil
	}
	argsNum = argsNum - 1

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
