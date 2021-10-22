package inmemory

import (
	"context"
	"reflect"
	"testing"
	"time"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/entities"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/errors"
	"github.com/stretchr/testify/assert"
)

func prepareStorage() *Storage {
	storage := New()
	storage.events = map[int]entities.Event{
		1: {
			ID:        1,
			Owner:     "Vlad",
			Title:     "task1",
			Text:      "text1",
			StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
		},
		2: {
			ID:        2,
			Owner:     "xd",
			Title:     "task2",
			Text:      "text2",
			StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 12, 22, 34, 15, 0, time.UTC),
		},
		3: {
			ID:        3,
			Owner:     "Vladimir",
			Title:     "task11",
			Text:      "text11",
			StartTime: time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 13, 21, 34, 15, 0, time.UTC),
		},
		4: {
			ID:        4,
			Owner:     "Test4",
			Title:     "xdxdsadfd",
			Text:      "asdfasdfasdf",
			StartTime: time.Date(2021, time.April, 15, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 16, 21, 34, 15, 0, time.UTC),
		},
	}
	return storage
}

func TestNew(t *testing.T) {
	storage1 := New()
	storage2 := New()
	if !reflect.DeepEqual(storage1, storage2) {
		t.Errorf("Not equal data in storage: %v, %v", storage1, storage2)
	}
}

func TestStorage_Add(t *testing.T) {
	storage := New()
	newEvent := entities.Event{
		ID:        4,
		Owner:     "Test4",
		Title:     "xdxdsadfd",
		Text:      "asdfasdfasdf",
		StartTime: time.Date(2021, time.April, 15, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 16, 21, 34, 15, 0, time.UTC),
	}
	err := storage.SaveEvent(context.Background(), newEvent)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	assert.Equal(t, len(storage.events), 1, "storage len should be 1")
}

func TestAddToNotEmptyStorage(t *testing.T) {
	storage := prepareStorage()
	newEvent := entities.Event{
		ID:        5,
		Owner:     "Test5",
		Title:     "xdxdsadfd",
		Text:      "asdfasdfasdf",
		StartTime: time.Date(2021, time.April, 17, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 18, 21, 34, 15, 0, time.UTC),
	}
	err := storage.SaveEvent(context.Background(), newEvent)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	assert.Equal(t, len(storage.events), 5, "len storage should be 5")
}

func TestAddEventToUsedDate(t *testing.T) {
	storage := prepareStorage()
	newEvent := entities.Event{
		ID:        5,
		Owner:     "Test4",
		Title:     "xdxdsadfd",
		Text:      "asdfasdfasdf",
		StartTime: time.Date(2021, time.April, 15, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 16, 21, 34, 15, 0, time.UTC),
	}
	err := storage.SaveEvent(context.Background(), newEvent)

	assert.Equal(t, err, errors.ErrDateBusy, "error should be ErrDateBusy")
}

func TestStorage_Delete(t *testing.T) {
	storage := prepareStorage()

	err := storage.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	err = storage.Delete(context.Background(), 2)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	assert.Equal(t, len(storage.events), 2, "storage len should be 2")
}

func TestStorage_Delete_Wrong(t *testing.T) {
	storage := prepareStorage()
	err := storage.Delete(context.Background(), 6)
	assert.Equal(t, err, errors.ErrNotFound, "expected error not found")
}

func TestStorage_Edit(t *testing.T) {
	storage := prepareStorage()

	event, err := storage.GetEventByID(context.Background(), 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	event.Title = "new Title"
	err = storage.Edit(context.Background(), event.ID, event)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	newEvent, err := storage.GetEventByID(context.Background(), 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(newEvent, event) {
		t.Errorf("not equal events: %v, %v", event, newEvent)
	}
}

func TestStorage_GetEvents(t *testing.T) {
	storage := prepareStorage()
	events, err := storage.GetEvents(context.Background())

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if len(events) != len(storage.events) {
		t.Errorf("expected %d events get %d", len(storage.events), len(events))
	}
	for _, e := range events {
		assert.Equal(t, storage.events[e.ID], e, "GetEvents returned elements incorrect")
	}
}

func TestStorage_GetEventByID_NotFound(t *testing.T) {
	storage := prepareStorage()
	_, err := storage.GetEventByID(context.Background(), 7)
	assert.Equal(t, err, errors.ErrNotFound, "should get ErrNotFound")
}

func TestStorage_GetEvents_Empty(t *testing.T) {
	storage := New()
	_, err := storage.GetEvents(context.Background())
	assert.Equal(t, err, errors.ErrNotFound, "should get ErrNotFound")
}

func Test_inTimeSpan(t *testing.T) {
	test := []struct {
		start  string
		end    string
		check  string
		isTrue bool
	}{
		{"23:00", "05:00", "04:00", true},
		{"23:00", "05:00", "23:30", true},
		{"23:00", "05:00", "20:00", false},
		{"10:00", "21:00", "11:00", true},
		{"10:00", "21:00", "22:00", false},
		{"10:00", "21:00", "03:00", false},
		{"22:00", "02:00", "00:00", true},
		{"10:00", "21:00", "10:00", true},
		{"10:00", "21:00", "21:00", true},
		{"23:00", "05:00", "06:00", false},
		{"23:00", "05:00", "23:00", true},
		{"23:00", "05:00", "05:00", true},
		{"10:00", "21:00", "10:00", true},
		{"10:00", "21:00", "21:00", true},
		{"10:00", "10:00", "09:00", false},
		{"10:00", "10:00", "11:00", false},
		{"10:00", "10:00", "10:00", true},
	}
	newLayout := "15:04"
	for _, row := range test {
		check, _ := time.Parse(newLayout, row.check)
		start, _ := time.Parse(newLayout, row.start)
		end, _ := time.Parse(newLayout, row.end)
		result := inTimeSpan(start, end, check)
		assert.Equal(t, result, row.isTrue, "Wrong inTimeSpan result")
	}
}
