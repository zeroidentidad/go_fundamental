package main

import (
	"../list"
	"../todo"
)

func main() {
	myToDos := factoryToDo("unique" /*"any"*/)
	rendering := factoryList("bullet")
	myToDos.SetList(rendering)

	myToDos.Add("Qué estudiar?")
	myToDos.Add("Explicarlo con palabras sencillas")
	myToDos.Add("Hacer con lo que aprendimos")
	myToDos.Add("Revisar y mejorar")
	myToDos.Add("Qué estudiar?")
	myToDos.Print()
}

func factoryToDo(s string) todo.ToDo {
	if s == "unique" {
		return todo.NewUnique()
	}

	return todo.NewAny()
}

func factoryList(s string) list.List {
	if s == "plain" {
		return list.NewPlain()
	}

	return list.NewBullet('*')
}
