package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

// File is responsbile for managing all endpoint handlers involving events

// Returns all events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events. Try again later."},
		)
		return
	}
	// Gin package automatically transforms data into JSON
	context.JSON(http.StatusOK, events)
}

// Returns an event belonging to specified id
func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param() allows us to get the value of the path parameter
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse event id."},
		)
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch event."},
		)
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// Bind user request to the above variable, event
	err := context.ShouldBindJSON(&event) // func needs a pointer to the object, event

	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Retrieve user Id from gin context; call this specific function
	// b/c it gives us the correct type for the event struct
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not create event. Try again later."},
		)
		return
	}

	// Send back OK status code and the event that was created
	context.JSON(
		http.StatusCreated,
		gin.H{"message": "Event created!", "event": event},
	)
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param() allows us to get the value of the path parameter
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse event id."},
		)
		return
	}

	// Look up id in db to see if it exists
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch the event."},
		)
		return
	}

	// Check if the user ID attached to the event matches the user ID we
	// pulled from this login token
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	// Create new event
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // Bind request data to updatedEvent

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse request data."},
		)
		return
	}

	// Give updatedEvent the previous event's ID
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not update event."},
		)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	// Get event ID from user request
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param() allows us to get the value of the path parameter
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse event id."},
		)
		return
	}

	userId := context.GetInt64("userId")
	// Look up id in db to see if it exists
	event, err := models.GetEventById(eventID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch the event."},
		)
		return
	}

	// Check if the user ID attached to the event matches the user ID we
	// pulled from this login token
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	// Delete the event
	err = event.Delete()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not delete the event."},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
