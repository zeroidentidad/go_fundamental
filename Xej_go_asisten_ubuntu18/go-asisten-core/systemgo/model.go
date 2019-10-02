package systemgo

import (
	"errors"
	"log"
	"reflect"
)

// SystemGoModel export
type SystemGoModel struct {
	FindPlayEvents chan string
}

// NewSystemGoModel func
func NewSystemGoModel() *SystemGoModel {
	return &SystemGoModel{
		FindPlayEvents: make(chan string),
	}
}

// FindRun func
func (ctx *SystemGoModel) FindRun(name string, values ...string) error {
	log.Printf("Name(%s) Params(%v) \n", name, values)
	salida := reflect.ValueOf(ctx).MethodByName(name)
	if salida.IsValid() == false {
		return errors.New("Is Not Valid")
	}

	inputs := make([]reflect.Value, len(values))
	for i := range values {
		inputs[i] = reflect.ValueOf(values[i])
	}
	salida.Call(inputs)
	return nil
}
