package modules

import (
	"fmt"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
)

// This module counts the number of distinct users who accessed the system
type UserAccessCount struct {
	UsersSeen map[string]bool
}

func (u *UserAccessCount) Setup() {
	u.UsersSeen = map[string]bool{}
}

func (u *UserAccessCount) Read(row models.Row) {
	u.UsersSeen[row.Username] = true
}

func (u *UserAccessCount) Display() string {
	return fmt.Sprintf("User access count: %d\n", len(u.UsersSeen))
}
