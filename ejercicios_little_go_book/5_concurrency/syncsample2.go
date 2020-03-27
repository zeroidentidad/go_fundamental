package main

import (
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func main() {
	go func() { lock.Lock() }()
	time.Sleep(time.Millisecond * 10)
	lock.Lock()
}
