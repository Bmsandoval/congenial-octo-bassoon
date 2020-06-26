package test

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

type uploadsOverFiftyTestData struct {
	Description string
	Input []models.Row
	Output string
}
func TestUploadsOverFiftyModule(t *testing.T) {
	testData := []uploadsOverFiftyTestData{
		{
			Description: "test single positive entry",
			Input: []models.Row {
				{
					 Size: 51,
					 Operation: "upload",
				},
			},
			Output: "Uploads over 50kb: 1\n",
		},
		{
			Description: "test one positive one negative",
			Input: []models.Row{
				{
					Size: 53,
					Operation: "upload",
				},
				{
					Size: 48,
					Operation: "upload",
				},
			},
			Output: "Uploads over 50kb: 1\n",
		},
		{
			Description: "test single negative entry",
			Input: []models.Row{
				{
					Size: 45,
					Operation: "upload",
				},
			},
			Output: "Uploads over 50kb: 0\n",
		},
		{
			Description: "test single download",
			Input: []models.Row{
				{
					Size: 55,
					Operation: "download",
				},
			},
			Output: "Uploads over 50kb: 0\n",
		},
	}

	for _, data := range testData {
		t.Run(data.Description, func(t *testing.T) {
			uploadsOverFifty := modules.UploadsOverFifty{}
			uploadsOverFifty.Setup()
			for _, row := range data.Input {
				uploadsOverFifty.Read(row)
			}
			responseData := uploadsOverFifty.Display()

			assert.Equal(t, data.Output, responseData)
		})
	}
}
