package modules

import "github.com/bmsandoval/congenial-octo-bassoon/internal/models"

type ModuleInterface interface {
	Read(row models.Row)
	Display() string
}