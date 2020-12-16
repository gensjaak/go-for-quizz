package controllers

import (
	"gensjaak/go-for-quizz/src/models"
	"gensjaak/go-for-quizz/src/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser <- Creates an user
func CreateUser(c *gin.Context) {
	// Fetch values for each field
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	userRoleID, _ := strconv.Atoi(c.PostForm("userRoleId"))

	// Pass values to the type and create the question
	status, err := models.CreateUser(types.UserAccount{
		ID:         0,
		FirstName:  firstName,
		LastName:   lastName,
		UserName:   userName,
		Password:   password,
		UserRoleID: userRoleID,
	})

	// Respond to the client according to the error status
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Created successfully",
			"data":    status,
		})
	}
}
