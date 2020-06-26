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
			Description: "test single user",
			Input: []models.Row{
				{
					Username: "harrypotter",
				},
			},
			Output: "User access count: 1\n",
		},
		{
			Description: "test multiple distinct users",
			Input: []models.Row{
				{
					Username: "harrypotter",
				},
				{
					Username: "ronweasley",
				},
			},
			Output: "User access count: 2\n",
		},
		{
			Description: "test duplicate user",
			Input: []models.Row{
				{
					Username: "harrypotter",
				},
				{
					Username: "ronweasley",
				},
				{
					Username: "harrypotter",
				},
			},
			Output: "User access count: 2\n",
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
