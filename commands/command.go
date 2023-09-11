package commands

import (
	"sync"
)

type Command struct {
	Name    string
	ArgsNum int
	Args    []string
}

var (
	dataStore = make(map[string]string)
	mutex     = &sync.Mutex{}
)
