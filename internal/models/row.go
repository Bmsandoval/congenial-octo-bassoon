package models

import "time"

type Row struct {
	Timestamp time.Time `csv:"timestamp"`
	Username string `csv:"username"`
	Operation string `csv:"operation"`
	Size int `csv:"size"`
}