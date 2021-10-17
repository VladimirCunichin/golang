package models

import (
	"github.com/vladimircunichin/golang/calendar/internal/types"
)

// StorageInterface interface
type StorageInterface interface {

	// Add types to storage.
	Add(event types.Event) (int64, error)

	// Edit types data in data storage
	Edit(id int64, event types.Event) error

	// GetEvents return all events
	GetEvents() ([]types.Event, error)

	//GetEventByID return types with ID
	GetEventByID(id int64) ([]types.Event, error)

	//Delete will mark types as deleted
	Delete(id int64) error
}
