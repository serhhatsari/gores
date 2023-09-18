package commands

import "strconv"

func rpop(key string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	list, ok := listStore[key]
	if !ok {
		return "", false
	}

	value, ok := list.RemoveLast()
	if !ok {
		return "", false
	}

	return value, true
}

func handleRPopCmd(command *Command) string {
	if command.ArgsNum != 1 && command.ArgsNum != 2 {
		return "-ERR wrong number of arguments for LPOP\r\n"
	}

	key := command.Args[0]
	count := 1
	if command.ArgsNum == 2 {
		var err error
		count, err = strconv.Atoi(command.Args[1])
		if err != nil {
			return "-ERR value is not an integer or out of range\r\n"
		}
	}

	var elements []string

	for i := 0; i < count; i++ {
		value, ok := rpop(key)
		if !ok {
			break
		}
		elements = append(elements, value)
	}

	if len(elements) == 0 {
		return "$-1\r\n"
	}

	// If there is only one element, return it as a bulk string
	if len(elements) == 1 {
		return "$" + strconv.Itoa(len(elements[0])) + "\r\n" + elements[0] + "\r\n"
	}

	// Otherwise, return it as an array
	res := "*" + strconv.Itoa(len(elements)) + "\r\n"
	for _, element := range elements {
		res += "$" + strconv.Itoa(len(element)) + "\r\n" + element + "\r\n"
	}
	return res
}
