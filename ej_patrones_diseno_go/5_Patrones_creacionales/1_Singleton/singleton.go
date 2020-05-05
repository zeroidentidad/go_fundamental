package singleton

import (
	"sync"
	"time"
)

// Singleton
type Singleton struct {
	Tiempo int64
}

// Creador "estático"
var instancia *Singleton
var once sync.Once

func GetInstancia() *Singleton {
	once.Do(func() {
		instancia = &Singleton{
			time.Now().Unix(),
		}
	})

	return instancia
}
