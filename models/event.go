package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

// Method to save an event; recall that to create a method you must add receive argument before function name
func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
