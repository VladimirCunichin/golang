package entities

import (
	"time"
)

type Event struct {
	ID        int
	Owner     string
	Title     string
	Text      string
	StartTime time.Time
	EndTime   time.Time
}
