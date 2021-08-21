package benchmark

import (
	"testing"

	v4 "github.com/google/uuid"
	v7 "github.com/uuid6/uuid6go-proto"
)

func GeneratorV4() string {

	id := v4.New().String()
	return id

}

func GeneratorV7() string {
	var gen v7.UUIDv7Generator
	id := gen.Next().ToString()
	return id
}

func BenchmarkUUIDv4(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		GeneratorV4()
	}
}

func BenchmarkV7(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		GeneratorV7()
	}
}
