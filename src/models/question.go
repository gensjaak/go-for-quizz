package models

import (
	"gensjaak/go-for-quizz/src/types"
)

// CreateTrueOrFalseQuestion <- Inserts true or false question
func CreateTrueOrFalseQuestion(data types.TrueOrFalseQuestion) (bool, error) {
	sql := `SELECT now()`

	_, err := Database.Exec(sql)

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateSliderQuestion <- Inserts slider question
func CreateSliderQuestion(data types.SliderQuestion) (bool, error) {
	sql := `SELECT now()`

	_, err := Database.Exec(sql)

	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateMultipleChoiceQuestion <- Inserts multiple choice question
func CreateMultipleChoiceQuestion(data types.MutltipleChoiceQuestion) (int, error) {
	sql := `SELECT now()`

	_, err := Database.Exec(sql)

	if err != nil {
		return 1, err
	}

	return 0, nil
}

// CreateMultipleChoiceQuestionAnswers <- Inserts multiple choice question proposals
func CreateMultipleChoiceQuestionAnswers(data []types.MutltipleChoiceQuestionAnswer) (bool, error) {
	sql := `SELECT now()`

	_, err := Database.Exec(sql)

	if err != nil {
		return false, err
	}

	return true, nil
}
