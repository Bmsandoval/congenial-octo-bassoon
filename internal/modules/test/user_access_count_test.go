package test

import (
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userAccessCountTestData struct {
	Description string
	Input []models.Row
	Output string
}
func TestUserAccessCountModule(t *testing.T) {
	testData := []userAccessCountTestData{
		{
			Description: "happy path",
			Input: []models.Row{
				{
					Username: "harry potter",
				},
			},
			Output: "User access count: 1\n",
		},
	}

	for _, data := range testData {
		t.Run(data.Description, func(t *testing.T) {
			userAccessCount := modules.UserAccessCount{}
			userAccessCount.Setup()
			for _, row := range data.Input {
				userAccessCount.Read(row)
			}
			responseData := userAccessCount.Display()

			assert.Equal(t, data.Output, responseData)
		})
	}
}
