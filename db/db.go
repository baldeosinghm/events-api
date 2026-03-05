package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
In the above imports, you may notice that while we import "github.com/mattn/go-sqlite3"
to make use of the sqlite3 package (GO's sql package requires a driver and we will
use this one), we also import the "database/sql" package (which is part of GO's
standard library).

The underhood functionality provided by sqlite3 is used by "database/sql" and so
that is why we import it.  We also place an underscore in front it b/c we don't
want Go to remove the import despite it not being called.
*/

// Create var, DB (capitalzied) to create a database connection pool that can be
// accessed by other packages
// Be wary of this code smell; may prefer dependency injection in future
var DB *sql.DB

// Initialize database (create db), open it for connection, and create tables
func InitDB() {
	// If data source is a file and it doesn't exist, it will automatically be created
	var err error

	// Open() opens the db to connection
	DB, err = sql.Open("sqlite3", "api.db") // dataSourceName: Path that leads to db

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10) // Sets max db connections
	DB.SetMaxIdleConns(5)

	createTables()
}

// Create tables so we can store data in db
func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	// Create a table that replicates the Event struct
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	// Execute database table
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
