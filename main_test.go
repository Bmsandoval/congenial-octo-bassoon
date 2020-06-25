package main

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal"
	"testing"
)

func BenchmarkParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		internal.ParseFile("sample_server_log.csv")
	}
}
