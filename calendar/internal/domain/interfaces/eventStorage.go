package interfaces

import (
	"context"
	"time"

	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *entities.Event) error
	GetEventByID(ctx context.Context, id string) (*entities.Event, error)
	GetEventByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*entities.Event
	GetAllEvents() []*entities.Event
}
