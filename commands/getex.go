package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func setExpire(key string, expire string, exprTime int) {
	var timer *time.Timer

	switch expire {
	case "EX":
		timer = time.NewTimer(time.Second * time.Duration(exprTime))
		break
	case "PX":
		timer = time.NewTimer(time.Millisecond * time.Duration(exprTime))
		break
	case "EXAT":
		durationUntilExpire := time.Unix(int64(exprTime), 0).Sub(time.Now())
		if durationUntilExpire < 0 {
			del(key)
			return
		}
		timer = time.NewTimer(durationUntilExpire)
		break
	case "PXAT":
		durationUntilExpire := time.Unix(0, int64(exprTime)*int64(time.Millisecond)).Sub(time.Now())
		if durationUntilExpire < 0 {
			del(key)
			return
		}
		timer = time.NewTimer(durationUntilExpire)
		break
	}
	<-timer.C

	del(key)
}

func handleGetExCmd(command *Command) string {
	if command.ArgsNum != 1 && command.ArgsNum != 3 {
		return "-ERR Wrong number of arguments for GET\r\n"
	}

	key := command.Args[0]
	value, ok := get(key)
	if !ok {
		return "$-1\r\n"
	}

	if command.ArgsNum == 3 {
		expire := command.Args[1]
		expire = strings.ToUpper(expire)
		if expire != "EX" && expire != "PX" && expire != "EXAT" && expire != "PXAT" {
			return "-ERR Syntax error\r\n"
		}

		exprTime, err := strconv.Atoi(command.Args[2])
		if err != nil {
			return "-ERR Syntax error\r\n"
		}
		go setExpire(key, expire, exprTime)
	}

	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}
