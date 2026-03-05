
# Event Booking API

The Event Booking API is a REST API that allows an individual to create a booking event (e.g. party, venue, trip, etc.).

***

## Features

- CRUD Operations - Create a user, login, and create, update, delete or even register for an event.

- User Authentication - Events can only be created, deleted, updated, or registered for by users who are logged in.  This authentication is achieved with JSON Web Tokens (JWT) which return a single token assosciated with the logged-in user.

- Database - User info, events, and event registrations are stored in a SQLite database.

- Route Testing - Route testing is done locally with help of REST Client, a VSCode extension.

- Middleware - To enforce user authentication, middleware is implemented to all routes requesting to make some change to an event.

## Tech Stack

- Language: Go
- Framework: Gin
- Database & Auth: SQLite3 and JWT
- HTTP Client - REST Client (Huachao Mao, VS Code)

## Quick Start

### Prerequisites

- Go 1.25+
- SQLite3
- HTTP Client - REST Client (Huachao Mao, VS Code)


### Installations

Clone the repository

```
git clone https://github.com/baldeosinghm/booking-events-api.git
```

Build the package

```
go build .
```

Run the executable

```
go run .
```

Feel free to edit the http test files so you can create, delete, update, or do as much or whatever your heart desires!