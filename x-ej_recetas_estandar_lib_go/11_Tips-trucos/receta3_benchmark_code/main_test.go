package main

import (
	"log"
	"testing"
)

// go test -run BenchmarkSampleStruct -v --bench . --benchmem
func BenchmarkSampleStruct(b *testing.B) {
	logger := log.New(devNull{}, "test", log.Llongfile)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Println("Is awesome")
	}
	b.StopTimer()
}

type devNull struct{}

func (d devNull) Write(b []byte) (int, error) {
	return 0, nil
}
