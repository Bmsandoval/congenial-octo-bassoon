package modules

import (
	"fmt"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
)

type UploadsOverFifty struct {
	Count int
}

func (u *UploadsOverFifty) Setup() {
	u.Count = 0
}

func (u *UploadsOverFifty) Read(row models.Row) {
	if row.Size > 50 {
		u.Count++
	}
}

func (u *UploadsOverFifty) Display() string {
	return fmt.Sprintf("Uploads over 50kb: %d\n", u.Count)
}
