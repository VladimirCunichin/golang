package inmemory

import (
	"context"
	"reflect"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/vladimircunichin/golang/calendar/internal/domain/entities"
	"github.com/vladimircunichin/golang/calendar/internal/domain/errors"
)

var (
	id1 = uuid.NewV4()
	id2 = uuid.NewV4()
	id3 = uuid.NewV4()
	id4 = uuid.NewV4()
)

func prepareStorage() *Storage {
	storage := New()
	storage.events = map[uuid.UUID]entities.Event{
		id1: {
			ID:        id1,
			Owner:     "Vlad",
			Title:     "task1",
			Text:      "text1",
			StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
		},
		id2: {
			ID:        id2,
			Owner:     "xd",
			Title:     "task2",
			Text:      "text2",
			StartTime: time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 12, 22, 34, 15, 0, time.UTC),
		},
		id3: {
			ID:        id3,
			Owner:     "Vladimir",
			Title:     "task11",
			Text:      "text11",
			StartTime: time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC),
			EndTime:   time.Date(2021, time.April, 13, 21, 34, 15, 0, time.UTC),
		},
		id4: {
			ID:        id4,
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
		ID:        uuid.NewV4(),
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
	if len(storage.events) != 1 {
		t.Errorf("events size is not 1")
	}
}

func TestAddToNotEmptyStorage(t *testing.T) {
	storage := prepareStorage()
	newEvent := entities.Event{
		ID:        uuid.NewV4(),
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
	if len(storage.events) != 5 {
		t.Errorf("events size is not 5")
	}
}

func TestAddEventToUsedDate(t *testing.T) {
	storage := prepareStorage()
	newEvent := entities.Event{
		ID:        uuid.NewV4(),
		Owner:     "Test4",
		Title:     "xdxdsadfd",
		Text:      "asdfasdfasdf",
		StartTime: time.Date(2021, time.April, 15, 21, 34, 15, 0, time.UTC),
		EndTime:   time.Date(2021, time.April, 16, 21, 34, 15, 0, time.UTC),
	}
	err := storage.SaveEvent(context.Background(), newEvent)
	if err == nil {
		t.Errorf("expected error: %s, but get nil", errors.ErrDateBusy)
	} else if err != errors.ErrDateBusy {
		t.Errorf("expected error: %s, get %s", errors.ErrDateBusy, err.Error())
	}
}

func TestStorage_Delete(t *testing.T) {
	storage := prepareStorage()
	keys := make([]uuid.UUID, 0, len(storage.events))
	for u := range storage.events {
		keys = append(keys, u)
	}

	err := storage.Delete(context.Background(), keys[1])
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	err = storage.Delete(context.Background(), keys[2])
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if len(storage.events) != 2 {
		t.Errorf("len should be 2, have %v", len(storage.events))
	}
}

func TestStorage_Delete_Wrong(t *testing.T) {
	storage := prepareStorage()
	err := storage.Delete(context.Background(), uuid.NewV4())
	if err == nil {
		t.Errorf("expected error: %s", errors.ErrNotFound)
	} else if err != errors.ErrNotFound {
		t.Errorf("expected error: %s, get : %s", errors.ErrNotFound, err.Error())
	}
}

func TestStorage_Edit(t *testing.T) {
	storage := prepareStorage()
	keys := make([]uuid.UUID, 0, len(storage.events))
	for u := range storage.events {
		keys = append(keys, u)
	}

	event, err := storage.GetEventByID(context.Background(), keys[0])
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	event.Title = "new Title"
	err = storage.Edit(context.Background(), event.ID, event)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	newEvent, err := storage.GetEventByID(context.Background(), keys[0])
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
		if storage.events[e.ID] != e {
			t.Errorf("events that were retrieved are wrong")
		}
	}
}

func TestStorage_GetEventByID_NotFound(t *testing.T) {
	storage := prepareStorage()
	_, err := storage.GetEventByID(context.Background(), uuid.NewV4())
	if err == nil {
		t.Errorf("expected error: %s, get nil", errors.ErrNotFound)
	} else if err != errors.ErrNotFound {
		t.Errorf("expected error: %s, get: %s", errors.ErrNotFound, err.Error())
	}
}

func TestStorage_GetEvents_Empty(t *testing.T) {
	storage := New()

	_, err := storage.GetEvents(context.Background())
	if err == nil {
		t.Errorf("expected error: %s, get %s", errors.ErrNotFound, err)
	}
	if err != errors.ErrNotFound {
		t.Errorf("expected error: %s, get: %s", errors.ErrNotFound, err.Error())
	}
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
		if result != row.isTrue {
			t.Errorf("get %t, expected %t on row: {%s, %s, %s, %t}", result, row.isTrue, row.start, row.end, row.check, row.isTrue)
		}
	}
}
