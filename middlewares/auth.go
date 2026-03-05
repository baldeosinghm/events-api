package middlewares

import (
	"net/http"

	"example.com/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

// Middleware to run during other handler requests
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Set() allows us to add some values to the gin context value
	// This is handy b/c we can then access values from outside functions w/o
	// having to return a variable from this function and calling it elsewhere
	context.Set("userId", userId)

	// Ensures the next handler request in line will execute correctly
	context.Next()
}
