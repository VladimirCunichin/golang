package inmemory

import (
	"context"
	"sync"
	"time"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/entities"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/errors"
)

// Storage struct
type Storage struct {
	mu     sync.Mutex
	events map[int]entities.Event
}

//New returns new storage
func New() *Storage {
	return &Storage{events: make(map[int]entities.Event)}
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
	storage.mu.Lock()
	storage.events[event.ID] = event
	storage.mu.Unlock()
	return nil
}

// Edit models data in data storage
func (storage *Storage) Edit(ctx context.Context, id int, event entities.Event) error {
	_, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	storage.mu.Lock()
	storage.events[id] = event
	storage.mu.Unlock()
	return nil
}

// GetEvents return all events
func (storage *Storage) GetEvents(ctx context.Context) ([]entities.Event, error) {
	if len(storage.events) > 0 {
		events := make([]entities.Event, 0, len(storage.events))
		storage.mu.Lock()
		for _, e := range storage.events {
			events = append(events, e)
		}
		storage.mu.Unlock()
		if len(events) > 0 {
			return events, nil
		}
	}
	return []entities.Event{}, errors.ErrNotFound
}

//GetEventByID return event by id
func (storage *Storage) GetEventByID(ctx context.Context, id int) (entities.Event, error) {
	storage.mu.Lock()
	e, ok := storage.events[id]
	storage.mu.Unlock()
	if !ok {
		return entities.Event{}, errors.ErrNotFound
	}
	return e, nil
}

//Delete will mark models as deleted
func (storage *Storage) Delete(ctx context.Context, id int) error {
	storage.mu.Lock()
	_, ok := storage.events[id]
	if !ok {
		return errors.ErrNotFound
	}
	delete(storage.events, id)
	storage.mu.Unlock()
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
