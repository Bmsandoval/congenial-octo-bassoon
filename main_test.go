package main

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"github.com/bmsandoval/congenial-octo-bassoon/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		userAccessCount := modules.UserAccessCount{}
		userAccessCount.Setup()


		uploadsOverFifty := modules.UploadsOverFifty{}
		uploadsOverFifty.Setup()

		jeffsUploadCount := modules.JeffsUploadCount{}
		jeffsUploadCount.Setup()

		result := pkg.CaptureOutput(internal.ParseFile, "sample_server_log_short.csv", &userAccessCount, &uploadsOverFifty, &jeffsUploadCount)
		assert.Equal(t, "User access count: 6\nUploads over 50kb: 8\nJeff's uploads on Apr 15th 2020: 2\n", result)
	})
}

func BenchmarkParseFile(b *testing.B) {
	userAccessCount := modules.UserAccessCount{}
	userAccessCount.Setup()

	uploadsOverFifty := modules.UploadsOverFifty{}
	uploadsOverFifty.Setup()

	jeffsUploadCount := modules.JeffsUploadCount{}
	jeffsUploadCount.Setup()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		internal.ParseFile("sample_server_log.csv", &userAccessCount, &uploadsOverFifty, &jeffsUploadCount)
	}
}
