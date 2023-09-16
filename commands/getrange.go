package commands

import "strconv"

func handleGetRangeCmd(command *Command) string {
	if command.ArgsNum != 3 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	start := command.Args[1]
	startVal, err := strconv.Atoi(start)
	if err != nil {
		return "-ERR Wrong type of argument\r\n"
	}

	end := command.Args[2]
	endVal, err := strconv.Atoi(end)
	if err != nil {
		return "-ERR Wrong type of argument\r\n"
	}

	val, ok := get(key)
	if !ok {
		return "+\"\"\r\n"
	}

	if startVal < 0 {
		startVal = len(val) + startVal
		if startVal < 0 {
			startVal = 0
		}
	}

	if endVal < 0 {
		endVal = len(val) + endVal
		if endVal < 0 {
			return "+\"\"\r\n"
		}
	}

	if startVal >= len(val) {
		return "+\"\"\r\n"
	}

	if endVal >= len(val) {
		endVal = len(val) - 1
	}

	if startVal > endVal {
		return "+\"\"\r\n"
	}

	result := val[startVal : endVal+1]

	return "$" + strconv.Itoa(len(result)) + "\r\n" + result + "\r\n"
}
