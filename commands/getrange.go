package commands

import "strconv"

func handleGetRangeCmd(command *Command) string {
	if command.ArgsNum != 3 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	start, err := strconv.Atoi(command.Args[1])
	if err != nil {
		return "-ERR Wrong type of argument\r\n"
	}

	end, err := strconv.Atoi(command.Args[2])
	if err != nil {
		return "-ERR Wrong type of argument\r\n"
	}

	val, ok := get(key)
	if !ok {
		return "+\"\"\r\n"
	}

	if start < 0 {
		start = len(val) + start
		if start < 0 {
			start = 0
		}
	}

	if end < 0 {
		end = len(val) + end
		if end < 0 {
			return "+\"\"\r\n"
		}
	}

	if start >= len(val) {
		return "+\"\"\r\n"
	}

	if end >= len(val) {
		end = len(val) - 1
	}

	if start > end {
		return "+\"\"\r\n"
	}

	result := val[start : end+1]

	return "$" + strconv.Itoa(len(result)) + "\r\n" + result + "\r\n"
}
