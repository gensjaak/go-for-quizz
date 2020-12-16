package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateQuizz <- should create a quiz from a label, a author and list of questions id
func CreateQuizz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    nil,
	})
}

// EditQuizz <- Given a quiz id, should be able to edit it by adding or removing some questions
func EditQuizz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    nil,
	})
}

// UserTriedToAnswerQuizz <- This is the function responsible for pausing or reprising an user quiz try
func UserTriedToAnswerQuizz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    nil,
	})
}
