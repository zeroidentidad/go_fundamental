package main

import "fmt"

type ConsoleLogger struct{ Name string }

func main() {
	ConsoleLogger{"Fulanito"}.Log("kepedo")
}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(l.Name, message)
}
