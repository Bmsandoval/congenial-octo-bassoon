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
				},
			},
			Output: "Uploads over 50kb: 1\n",
		},
		{
			Description: "test one positive one negative",
			Input: []models.Row{
				{
					Size: 53,
				},
				{
					Size: 48,
				},
			},
			Output: "Uploads over 50kb: 1\n",
		},
		{
			Description: "test single negative entry",
			Input: []models.Row{
				{
					Size: 45,
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
