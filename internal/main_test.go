package main

import (
	"testing"
)

func BenchmarkAcquireServers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFile()
	}
}