package modules

import (
	"fmt"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
)

// This module will count number of times Jeff uploaded on Apr 15th 2020
type JeffsUploadCount struct {
	Count int
}

func (u *JeffsUploadCount) Setup() {
	u.Count = 0
}

func (u *JeffsUploadCount) Read(row models.Row) {
	if row.Timestamp.Day() == 15 && row.Timestamp.Month() == 4 && row.Timestamp.Year() == 2020 &&
		row.Username == "jeff22" &&
		row.Operation == "upload"{
		u.Count++
	}
}

func (u *JeffsUploadCount) Display() string {
	return fmt.Sprintf("Jeff's uploads on Apr 15th 2020: %d\n", u.Count)
}
