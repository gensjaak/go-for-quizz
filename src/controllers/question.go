package controllers

import (
	"gensjaak/go-for-quizz/src/models"
	"gensjaak/go-for-quizz/src/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create True or False question
func createTrueOrFalseQuestion(c *gin.Context) {
	// Get the label from the Multipar-Encoded form from the request
	label := c.PostForm("label")

	// Question type here is "TF" <- True or False
	questionType := "TF"

	// The answer is (bool | TRUE or FALSE)
	answerIs, _ := strconv.ParseBool(c.PostForm("answerIs"))

	// Pass values to the type and create the question
	status, err := models.CreateTrueOrFalseQuestion(types.TrueOrFalseQuestion{
		QuestionID:   0,
		Label:        label,
		QuestionType: questionType,
		AnswerIs:     answerIs,
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

// Create a slider question
func createSliderQuestion(c *gin.Context) {
	// Get the label from the Multipar-Encoded form from the request
	label := c.PostForm("label")

	// Question type here is "SL" <- Slider
	questionType := "SL"

	// Bounds of the correct answer
	answerMin, _ := strconv.Atoi(c.PostForm("answerMin"))
	answerMax, _ := strconv.Atoi(c.PostForm("answerMax"))

	// Pass values to the type and create the question
	status, err := models.CreateSliderQuestion(types.SliderQuestion{
		QuestionID:   0,
		Label:        label,
		QuestionType: questionType,
		AnswerMin:    answerMin,
		AnswerMax:    answerMax,
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

// Create multiple choice question
func createMultipleChoiceQuestion(c *gin.Context) {
	// Get the label from the Multipar-Encoded form from the request
	label := c.PostForm("label")

	// Question type here is "MC" <- Multiple Choice
	questionType := "MC"

	// At this version of the API, we only support 3 proposals for a multiple question
	// Here we are getting the proposals labels and their state
	proposal1 := c.PostForm("proposal1")
	proposal1IsCorrect, _ := strconv.ParseBool(c.PostForm("proposal1IsCorrect"))

	proposal2 := c.PostForm("proposal2")
	proposal2IsCorrect, _ := strconv.ParseBool(c.PostForm("proposal2IsCorrect"))

	proposal3 := c.PostForm("proposal3")
	proposal3IsCorrect, _ := strconv.ParseBool(c.PostForm("proposal3IsCorrect"))

	// Pass values to the type and create the question
	createdQuestionID, err := models.CreateMultipleChoiceQuestion(types.MutltipleChoiceQuestion{
		QuestionID:   0,
		Label:        label,
		QuestionType: questionType,
	})

	// Create the question proposals
	proposals := []types.MutltipleChoiceQuestionAnswer{
		types.MutltipleChoiceQuestionAnswer{
			MutltipleQuestionID: createdQuestionID,
			Label:               proposal1,
			IsCorrect:           proposal1IsCorrect,
		},
		types.MutltipleChoiceQuestionAnswer{
			MutltipleQuestionID: createdQuestionID,
			Label:               proposal2,
			IsCorrect:           proposal2IsCorrect,
		},
		types.MutltipleChoiceQuestionAnswer{
			MutltipleQuestionID: createdQuestionID,
			Label:               proposal3,
			IsCorrect:           proposal3IsCorrect,
		},
	}

	// Respond to the client according to the error status
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		status, err := models.CreateMultipleChoiceQuestionAnswers(proposals)

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
}

// CreateQuestion <- main function for creating a question
func CreateQuestion(c *gin.Context) {
	// Get the question type from the form and exec the related function
	questionType := c.PostForm("questionType")

	if questionType == "TF" {
		createTrueOrFalseQuestion(c)
	} else if questionType == "SL" {
		createSliderQuestion(c)
	} else if questionType == "MC" {
		createMultipleChoiceQuestion(c)
	}
}
