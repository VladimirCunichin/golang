package errors

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

const (
	ErrDateBusy      EventError = "this time is already in use"
	ErrNotFound      EventError = "event not found"
	ErrEventDeleted  EventError = "event was deleted"
	ErrEventIdExists EventError = "event with this ID already exists"
)
