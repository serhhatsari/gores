package commands

import "strconv"

func ltrim(key string, start, end int) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := listStore[key]; !ok {
		return
	}
	listStore[key].Trim(start, end)
}

func handleLTrimCmd(command *Command) string {
	if len(command.Args) != 3 {
		return "-ERR wrong number of arguments for LTRIM\r\n"
	}

	key := command.Args[0]
	start, err := strconv.Atoi(command.Args[1])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}

	end, err := strconv.Atoi(command.Args[2])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}

	ltrim(key, start, end)

	return "+OK\r\n"
}
