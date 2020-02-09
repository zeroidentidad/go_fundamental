package todo

import "../list"

// ToDo interface que tiene la lista de tareas
type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
