package usecases

import (
	"context"
	"log"
	"testing"
	"time"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/adapters/inmemory"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

var preparedUseCases *EventUseCases = prepareUseCases()

func prepareUseCases() *EventUseCases {
	usecases := New(inmemory.New())

	events := []entities.Event{
		{
			ID:        1,
			Owner:     "Vlad",
			Title:     "task1",
			Text:      "text1",
			StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
		},
		{
			ID:        2,
			Owner:     "xd",
			Title:     "task2",
			Text:      "text2",
			StartTime: time.Date(2021, time.April, 12, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 13, 22, 34, 15, 0, time.UTC),
		},
		{
			ID:        3,
			Owner:     "Vladimir",
			Title:     "task11",
			Text:      "text11",
			StartTime: time.Date(2021, time.April, 14, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 15, 21, 34, 15, 0, time.UTC),
		},
		{
			ID:        4,
			Owner:     "Test4",
			Title:     "xdxdsadfd",
			Text:      "asdfasdfasdf",
			StartTime: time.Date(2021, time.April, 17, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 18, 21, 34, 15, 0, time.UTC),
		},
	}
	for _, event := range events {
		_, err := usecases.SaveEvent(context.Background(), event.Owner, event.Title, event.Text, event.StartTime, event.EndTime)
		if err != nil {
			log.Fatalf("unexpected error: %s %v", err, event)
		}
	}
	return usecases
}

func TestGetEventByID(t *testing.T) {
	usecases := preparedUseCases
	testEvent := entities.Event{

		ID:        1,
		Owner:     "Vlad",
		Title:     "task1",
		Text:      "text1",
		StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
	}
	event, err := usecases.GetEventByID(context.Background(), 1)
	if err != nil {
		t.Errorf("GetEventByID error: %s", err)
	}
	assert.Equal(t, testEvent, event, "events should be equal")
}

func TestGetEvents(t *testing.T) {
	usecases := preparedUseCases
	events, err := usecases.GetEvents(context.Background())
	if err != nil {
		t.Errorf("getevents error: %s", err)
	}
	assert.Equal(t, 4, len(events), "events length should be 4")
}

func TestDelete(t *testing.T) {
	usecases := preparedUseCases
	err := usecases.Delete(context.Background(), 4)
	if err != nil {
		t.Errorf("delete error: %s", err)
	}
	events, err := usecases.GetEvents(context.Background())
	if err != nil {
		t.Errorf("getevents error: %s", err)
	}
	assert.Equal(t, 3, len(events), "events length should be 3")
}

func TestEdit(t *testing.T) {
	usecases := preparedUseCases
	editedEvent := entities.Event{
		ID:        2,
		Owner:     "editedOwner",
		Title:     "editedTitle",
		Text:      "editedText",
		StartTime: time.Date(2021, time.April, 22, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 23, 21, 34, 15, 0, time.UTC),
	}
	err := usecases.Edit(context.Background(), editedEvent.ID, editedEvent.Owner, editedEvent.Title, editedEvent.Text, editedEvent.StartTime, editedEvent.EndTime)
	if err != nil {
		t.Errorf("error during edit: %s", err)
	}
	resultEvent, err := usecases.GetEventByID(context.Background(), editedEvent.ID)
	if err != nil {
		t.Errorf("error during get event: %s", err)
	}
	assert.Equal(t, editedEvent, resultEvent, "edited event retrieved don't match to new event")
}

func TestSaveEvent(t *testing.T) {
	usecases := New(inmemory.New())

	expected, err := usecases.SaveEvent(context.Background(), "owner", "title", "text", time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC), time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC))
	if err != nil {
		t.Errorf("save event error %s", err)
	}
	getEvent, err := usecases.EventStorage.GetEventByID(context.Background(), expected.ID)
	if err != nil {
		t.Errorf("get eventbyid error %s", err)
	}
	assert.Equal(t, expected, getEvent, "events should be the same")
}
