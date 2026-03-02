package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// This is a simple Go REST API server buit w/ the Gin web framework

// Let's start by first handling an incoming request w/ gin
func main() {
	// Create Gin HTTP server w/ default middleware (logging + recovery)
	server := gin.Default()

	server.GET("/events", getEvents) // Create "/events" endpoint; make getEvents the endpoint handler
	server.POST("/events", createEvent)

	server.Run(":8080") // Run server on port 8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// Gin package automatically transforms data into JSON
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// Bind user request to the above event variable
	err := context.ShouldBindJSON(&event) // func needs a pointer to the object, event

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Create temporary dummy IDs
	event.ID = 1
	event.UserID = 1

	event.Save()

	// Send back OK status code and the event that was created
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
