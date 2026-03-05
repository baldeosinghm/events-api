package models

import (
	"time"

	"example.com/rest-api/db"
)

// File is responsbile for backend logic related to event(s)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// Method to save an event; recall that a method needs add a receiver argument before the function name
func (e *Event) Save() error {
	// Insert into db table, events, the field names you want populated
	// The question marks allow for a SQL injection, safe way of inserting
	// values into this query
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	// Prepare() prepares a SQL statement. Alternatively, you could also directly
	// execute a statement via Exec(). But Prepare() can lead to better performance
	// in certain situations.
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.UserID,
	)

	if err != nil {
		return err
	}
	// Get the id of the event that was inserted; the id is returned by the db
	id, err := result.LastInsertId()
	e.ID = id // id is of type int64 so update the Event struct ID type so it can hold that value

	return err
}

// Fetch events from table, events
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" // the query to select all rows from table, "events"
	rows, err := db.DB.Query(query) // store rows
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?" // This -> "?" allows for safe sql insertions
	row := db.DB.QueryRow(query, id)             // We can safely insert the value for "id" here

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return nil, err // The zero value for Event is an empty struct; you could use a pointer instead
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.ID,
	)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
