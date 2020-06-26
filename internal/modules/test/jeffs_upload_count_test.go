package test

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type jeffsUploadCountTestData struct {
	Description string
	Input []models.Row
	Output string
}
func TestJeffsUploadCountModule(t *testing.T) {
	aprFifteenth, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", "Wed Apr 15 05:14:41 UTC 2020")
	testData := []jeffsUploadCountTestData{
		{
			Description: "test wrong date",
			Input: []models.Row{
				{
					Username: "jeff22",
					Operation: "upload",
					Timestamp: time.Now(),
				},
			},
			Output: "Jeff's uploads on Apr 15th 2020: 0\n",
		},
		{
			Description: "test single positive event",
			Input: []models.Row{
				{
					Username: "jeff22",
					Operation: "upload",
					Timestamp: aprFifteenth,
				},
			},
			Output: "Jeff's uploads on Apr 15th 2020: 1\n",
		},
		{
			Description: "test jeff download",
			Input: []models.Row{
				{
					Username: "jeff22",
					Operation: "download",
					Timestamp: aprFifteenth,
				},
			},
			Output: "Jeff's uploads on Apr 15th 2020: 0\n",
		},
		{
			Description: "test non-jeff user",
			Input: []models.Row{
				{
					Username: "sarah",
					Operation: "upload",
					Timestamp: aprFifteenth,
				},
			},
			Output: "Jeff's uploads on Apr 15th 2020: 0\n",
		},
	}

	for _, data := range testData {
		t.Run(data.Description, func(t *testing.T) {
			jeffsUploadCount := modules.JeffsUploadCount{}
			jeffsUploadCount.Setup()
			for _, row := range data.Input {
				jeffsUploadCount.Read(row)
			}
			responseData := jeffsUploadCount.Display()

			assert.Equal(t, data.Output, responseData)
		})
	}
}
