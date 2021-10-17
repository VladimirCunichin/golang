package calendar

import (
	"time"

	"github.com/vladimircunichin/golang/calendar/internal/models"
	"github.com/vladimircunichin/golang/calendar/internal/types"
)

type Calendar struct {
	storage models.StorageInterface
}

//New will create new calendar
func New(storage models.StorageInterface) *Calendar {
	return &Calendar{storage: storage}
}

// Add event to calendar.
func (c Calendar) Add(title string, dateStarted time.Time, dateComplete time.Time, notice string, userId int64) (int64, error) {
	event := types.Event{
		ID:           0,
		Title:        title,
		DateEdited:   time.Now(),
		EditorID:     userId,
		DateCreated:  time.Now(),
		CreatorID:    userId,
		DateStarted:  dateStarted,
		DateComplete: dateComplete,
		Notice:       notice,
		Deleted:      false,
	}
	return c.storage.Add(event)
}

//Edit event data in calendar
func (c Calendar) Edit(id int64, event types.Event, userId int64) error {
	event.EditorID = userId
	event.DateEdited = time.Now()
	return c.storage.Edit(id, event)
}

//GetEvents return all events
func (c Calendar) GetEvents() ([]types.Event, error) {
	return c.storage.GetEvents()
}

//GetEventByID return event with ID
func (c Calendar) GetEventByID(id int64) ([]types.Event, error) {
	return c.storage.GetEventByID(id)
}

//Delete will mark event as deleted
func (c Calendar) Delete(id int64) error {
	return c.storage.Delete(id)
}
