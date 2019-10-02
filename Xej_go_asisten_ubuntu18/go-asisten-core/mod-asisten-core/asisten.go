package core

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

// WordsToAsistenResponse convert Words In Asisten Response
func (ctx *ModelCore) WordsToAsistenResponse(words string) (string, func()) {
	resp := ctx.Invoke("WordsToAsistenResponse", words)
	return resp[0].String(), resp[1].Interface().(func())
}

// WordsToSpeechBase64 convert Words In speech base64
func (ctx *ModelCore) WordsToSpeechBase64(words string) string {
	return ctx.Invoke("WordsToSpeechBase64", words)[0].String()
}

// Invoke methods by name
func (ctx *ModelCore) Invoke(name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(ctx.AsistenCore).MethodByName(name).Call(inputs)
}

// GetDB conection
func (ctx *ModelCore) GetDB() *gorm.DB {
	return reflect.ValueOf(ctx.AsistenCore).Elem().FieldByName("Db").Interface().(*gorm.DB)
}
