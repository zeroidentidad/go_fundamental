package singleton

import (
	"sync"
)

var (
	p    *person
	once sync.Once
)

func init() {
	once.Do(func() {
		p = &person{}
	})
}

/*
sobre once sync.Once puntualización: que no es para "todo el paquete". Es para la instancia de la variable sobre la que se está mandando el mensaje Do() [que en tu ejemplo es solo una, y se llama "once" la variable]. Eso quiere decir que si tuvieras 2 variables: once1 y once2 podrías ejecutar Do() una vez en cada una de ellas, así que podrías ejecutar 2 veces la invocación a Do(). Y así, lo mismo, si tuvieras N variables de tipo "sync.Once" podrías hacer 1 invocación por cada una de ellas [por lo tanto, N invocaciones]. Según la documentación: "Do calls the function f if and only if Do is being called for the first time for this instance of Once.". Lo importante ahí es el "this instance of Once".
*/

func GetInstance() *person {
	/* se paso a init
	once.Do(func() {
		p = &person{}
	})*/

	return p
}

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
