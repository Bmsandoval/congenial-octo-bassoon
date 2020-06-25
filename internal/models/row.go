package models

import "time"

type Row struct {
	Timestamp time.Time
	Username string
	Operation string
	Size int
}