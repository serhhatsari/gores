package commands

import (
	"fmt"
	"strconv"
)

func handleIncrByFloatCmd(command *Command) string {
	if command.ArgsNum != 2 {
		return "-ERR Wrong number of arguments for Command\r\n"
	}

	key := command.Args[0]
	value, ok := get(key)
	if !ok {
		set(key, "0")
	}

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "-ERR value is not an float or out of range\r\n"
	}

	incrBy, err := strconv.ParseFloat(command.Args[1], 64)
	if err != nil {
		return "-ERR value is not an float or out of range\r\n"
	}

	f += incrBy

	set(key, fmt.Sprintf("%f", f))

	return fmt.Sprintf("$%d\r\n%s\r\n", len(fmt.Sprintf("%f", f)), fmt.Sprintf("%f", f))
}
