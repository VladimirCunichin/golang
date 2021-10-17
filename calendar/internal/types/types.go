package types

import "time"

type Event struct {
	ID           int64
	Title        string
	DateEdited   time.Time
	EditorID     int64
	DateCreated  time.Time
	CreatorID    int64
	DateStarted  time.Time
	DateComplete time.Time
	Notice       string
	Deleted      bool
}
