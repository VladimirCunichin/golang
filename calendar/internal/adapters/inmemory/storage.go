package storage

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
	"github.com/vladimircunichin/golang/calendar/internal/domain/errors"
)

type EventUsecaseInterface interface {
	CreateEvent(ctx context.Context, owner, title, text string, startTime, endTime *time.Time) (*entities.Event, error)
}

// Storage struct
type Storage struct {
	EventUsecase EventUsecaseInterface
	events       map[uuid.UUID]entities.Event
}

//New returns new storage
func New() *Storage {
	return &Storage{events: make(map[uuid.UUID]entities.Event)}
}

// Add models to storage.
func (storage *Storage) Add(event entities.Event) (uuid.UUID, error) {
	// for _, e := range storage.events {
	// 	if inTimeSpan(*e.StartTime, e.DateComplete, event.DateStarted) ||
	// 		inTimeSpan(e.StartTime, e.DateComplete, event.DateComplete) ||
	// 		inTimeSpan(event.StartTime, event.DateComplete, e.DateStarted) ||
	// 		inTimeSpan(event.StartTime, event.DateComplete, e.DateComplete) {
	// 		return uuid.UUID{}, errors.ErrDateBusy
	// 	}
	// }

	_, ok := storage.events[event.ID]
	if ok {
		return uuid.UUID{}, errors.ErrEventIdExists
	}
	storage.events[event.ID] = event
	return event.ID, nil
}

// Edit models data in data storage
func (storage *Storage) Edit(id uuid.UUID, event entities.Event) error {
	_, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	storage.events[event.ID] = event
	return nil
}

// GetEvents return all events
func (storage *Storage) GetEvents() ([]entities.Event, error) {
	if len(storage.events) > 0 {
		events := make([]entities.Event, 0, len(storage.events))
		for _, e := range storage.events {
			events = append(events, e)
		}
		if len(events) > 0 {
			return events, nil
		}
	}
	return []entities.Event{}, errors.ErrNotFound
}

//GetEventByID return models with ID
func (storage *Storage) GetEventByID(id uuid.UUID) ([]entities.Event, error) {
	e, ok := storage.events[id]
	if !ok {
		return []entities.Event{}, errors.ErrNotFound
	}
	return []entities.Event{e}, nil
}

//Delete will mark models as deleted
func (storage *Storage) Delete(id uuid.UUID) error {
	e, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	storage.events[id] = e
	return nil
}

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
