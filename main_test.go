package main

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal"
	"github.com/bmsandoval/congenial-octo-bassoon/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result := pkg.CaptureOutput(internal.ParseFile, "sample_server_log_short.csv")
		assert.Equal(t, "User access count: 5", result)
	})
}

func BenchmarkParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		internal.ParseFile("sample_server_log.csv")
	}
}
