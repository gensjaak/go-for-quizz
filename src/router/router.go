package router

import (
	"gensjaak/go-for-quizz/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Just return a simple message showing that API was successfull started
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
		"data":    nil,
	})
}

// Route not configured
func routeNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Route not found on API",
		"success": false,
		"data":    nil,
	})
}

// Route not configured
func methodNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Method with that route not found on API",
		"success": false,
		"data":    nil,
	})
}

// AuthRequired <- should be checking if user is connected and get its role
func AuthRequired(c *gin.Context) {
	c.Next()
}

// Configure configure routes and handlers and return it
func Configure() *gin.Engine {
	router := gin.New()

	// Basic middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Authentification middleware and API grouping
	api := router.Group("/api/v1").Use(AuthRequired)
	{
		// Check API status
		api.GET("/", healthCheck)

		// Create any type of question
		api.POST("/questions/create", controllers.CreateQuestion)

		// Create a quiz
		api.POST("/quizzs/create", controllers.CreateQuizz)

		// Edit a quiz
		api.PATCH("/quizzs/edit/:id", controllers.EditQuizz)

		// Create an user
		api.POST("/users/create", controllers.CreateUser)

		// Respond to a quizz
		api.PATCH("/tries/:userId/:quizId", controllers.UserTriedToAnswerQuizz)
	}

	// When the route does not match any previous routes
	router.NoRoute(routeNotFound)

	// When the route method does not match any previous methods
	router.NoMethod(methodNotFound)

	// Router setup, return it
	return router
}
