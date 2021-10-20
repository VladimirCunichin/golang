package inmemory

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
	"github.com/vladimircunichin/golang/calendar/internal/domain/errors"
)

// Storage struct
type Storage struct {
	events map[uuid.UUID]entities.Event
}

//New returns new storage
func New() *Storage {
	return &Storage{events: make(map[uuid.UUID]entities.Event)}
}

// Add models to storage.
func (storage *Storage) SaveEvent(ctx context.Context, event entities.Event) error {
	for _, e := range storage.events {
		if inTimeSpan(e.StartTime, e.EndTime, event.StartTime) ||
			inTimeSpan(e.StartTime, e.EndTime, event.EndTime) {
			return errors.ErrDateBusy
		}
	}

	_, ok := storage.events[event.ID]
	if ok {
		return errors.ErrEventIdExists
	}
	storage.events[event.ID] = event
	return nil
}

// Edit models data in data storage
func (storage *Storage) Edit(ctx context.Context, id uuid.UUID, event entities.Event) error {
	_, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	storage.events[event.ID] = event
	return nil
}

// GetEvents return all events
func (storage *Storage) GetEvents(ctx context.Context) ([]entities.Event, error) {
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

//GetEventByID return event by id
func (storage *Storage) GetEventByID(ctx context.Context, id uuid.UUID) (entities.Event, error) {
	e, ok := storage.events[id]
	if !ok {
		return entities.Event{}, errors.ErrNotFound
	}
	return e, nil
}

//Delete will mark models as deleted
func (storage *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	_, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	delete(storage.events, id)
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
