package modules

import "github.com/bmsandoval/congenial-octo-bassoon/internal/models"

type ModuleInterface interface {
	Setup()
	Read(row models.Row)
	Display() string
}