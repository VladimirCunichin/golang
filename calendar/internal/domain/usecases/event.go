package usecases

import (
	"context"
	"time"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/entities"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/interfaces"
)

var currentId int = 0

type EventUseCases struct {
	EventStorage interfaces.EventStorage
}

func New(storage interfaces.EventStorage) *EventUseCases {
	return &EventUseCases{EventStorage: storage}
}

func (usecase *EventUseCases) SaveEvent(ctx context.Context, owner, title, text string, startTime, endTime time.Time) (entities.Event, error) {
	currentId++
	newId := currentId
	event := entities.Event{
		ID:        newId,
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := usecase.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (es *EventUseCases) GetEventByID(ctx context.Context, id int) (entities.Event, error) {
	return es.EventStorage.GetEventByID(ctx, id)
}
func (es *EventUseCases) GetEvents(ctx context.Context) ([]entities.Event, error) {
	return es.EventStorage.GetEvents(ctx)
}
func (es *EventUseCases) Delete(ctx context.Context, id int) error {
	return es.EventStorage.Delete(ctx, id)
}
func (es *EventUseCases) Edit(ctx context.Context, id int, owner, title, text string, startTime, endTime time.Time) error {
	event := entities.Event{
		ID:        id,
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	return es.EventStorage.Edit(ctx, id, event)
}
