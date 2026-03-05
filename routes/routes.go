package routes

import (
	"example.com/event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

// Responsible for registering event routes

func RegisterRoutes(server *gin.Engine) {
	// User Gin HTTP server to register routes
	// If you want to call multiple functions in a request, you can do so with Gin
	// It would look like this ->	server.POST("/events", middlewares.Authenticate, createEvent)
	server.GET("/events", getEvents) // Create "/events" endpoint; make getEvents the endpoint handler
	server.GET("/events/:id", getEvent)

	// The Group() method allows you to simplify the task of running something like
	// middleware before a bunch of other functions
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Use() adds middleware to the group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
