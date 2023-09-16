package commands

import "strconv"

func handleSetRangeCmd(command *Command) string {
	if command.ArgsNum != 3 {
		return "-ERR Wrong number of arguments\r\n"
	}

	key := command.Args[0]

	offset, err := strconv.Atoi(command.Args[1])
	if err != nil {
		return "-ERR Wrong offset\r\n"
	}

	changedVal := command.Args[2]

	val, ok := get(key)
	if !ok {
		val = changedVal
		set(key, val)
		return ":" + strconv.Itoa(len(val)) + "\r\n"
	}

	if offset < 0 {
		offset = len(val) + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset > len(val) {
		val += changedVal
		set(key, val)
		return ":" + strconv.Itoa(len(val)) + "\r\n"
	}

	if offset+len(changedVal) > len(val) {
		val = val[:offset] + changedVal
		set(key, val)
		return ":" + strconv.Itoa(len(val)) + "\r\n"
	}

	val = val[:offset] + changedVal + val[offset+len(changedVal):]
	set(key, val)
	return ":" + strconv.Itoa(len(val)) + "\r\n"
}
