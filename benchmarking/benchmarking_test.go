package benchmarking

import (
	"testing"
	"time"
)

// go test -bench=.go

func targetMethod() {
	for i := 0; i < 1000; i++ {
		duration := time.Duration(500) * time.Nanosecond
		time.Sleep(duration)
	}
}

func BenchmarkTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		targetMethod()
	}
}
