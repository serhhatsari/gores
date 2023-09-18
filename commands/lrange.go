package commands

import "strconv"

func lrange(key string, start int, end int) []string {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := listStore[key]; !ok {
		return nil
	}
	return listStore[key].Range(start, end)
}

func handleLRangeCmd(command *Command) string {
	if len(command.Args) != 3 {
		return "-ERR wrong number of arguments for LRange\r\n"
	}

	key := command.Args[0]
	start, err := strconv.Atoi(command.Args[1])
	if err != nil {
		return "-ERR Value is not an integer or out of range\r\n"
	}
	end, err := strconv.Atoi(command.Args[2])
	if err != nil {
		return "-ERR Value is not an integer or out of range\r\n"
	}

	// Check if key exists
	ok := exists(key)
	if !ok {
		return "*0\r\n"
	}

	result := lrange(key, start, end)
	if result == nil {
		return "*0\r\n"
	}

	response := "*" + strconv.Itoa(len(result)) + "\r\n"

	for _, value := range result {
		response += "$" + strconv.Itoa(len(value)) + "\r\n" + value + "\r\n"
	}

	return response
}
