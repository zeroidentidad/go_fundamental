package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	value      int
	executedBy string
}

var total int
var task Task
var lock sync.Mutex

func main() {

	fmt.Printf("synchronizing goroutines demo\n")
	total = 0
	task.value = 0
	task.executedBy = ""
	display()

	// run background
	go calculate()
	go perform()

	// press ENTER to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func calculate() {

	for total < 10 {
		lock.Lock()
		task.value = rand.Intn(100)
		task.executedBy = "from calculate()"
		display()
		total++
		lock.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func perform() {

	for total < 10 {
		lock.Lock()
		task.value = rand.Intn(100)
		task.executedBy = "from perform()"
		display()
		total++
		lock.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func display() {
	fmt.Println("--------------------------")
	fmt.Println(task.value)
	fmt.Println(task.executedBy)
	fmt.Println("--------------------------")

}
