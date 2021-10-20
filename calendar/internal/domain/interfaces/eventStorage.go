package interfaces

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event entities.Event) error
	GetEventByID(ctx context.Context, id uuid.UUID) (entities.Event, error)
	GetEvents(ctx context.Context) ([]entities.Event, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Edit(ctx context.Context, id uuid.UUID, event entities.Event) error
	// GetEventByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*entities.Event
}
