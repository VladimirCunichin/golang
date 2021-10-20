package usecases

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
	"github.com/vladimircunichin/golang/calendar/internal/domain/interfaces"
)

type EventUseCases struct {
	EventStorage interfaces.EventStorage
}

func (usecase *EventUseCases) CreateEvent(ctx context.Context, owner, title, text string, startTime, endTime time.Time) (*entities.Event, error) {
	event := &entities.Event{
		ID:        uuid.NewV4(),
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := usecase.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventUseCases) GetEventByID(ctx context.Context, id string) (*entities.Event, error) {
	panic("implement me")
}

func (es *EventUseCases) GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) ([]entities.Event, error) {
	panic("implement me")
}
