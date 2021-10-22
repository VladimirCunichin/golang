package interfaces

import (
	"context"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/entities"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event entities.Event) error
	GetEventByID(ctx context.Context, id int) (entities.Event, error)
	GetEvents(ctx context.Context) ([]entities.Event, error)
	Delete(ctx context.Context, id int) error
	Edit(ctx context.Context, id int, event entities.Event) error
	// GetEventByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*entities.Event
}
