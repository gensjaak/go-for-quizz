package types

// UserAccountRole ...
type UserAccountRole struct {
	ID      int
	Label   string
	IsAdmin bool
}

// UserAccount ...
type UserAccount struct {
	ID         int
	FirstName  string
	LastName   string
	UserName   string
	Password   string
	UserRoleID int
}

// Question ...
type Question struct {
	ID           int
	Label        string
	QuestionType string
}

// MutltipleChoiceQuestion ...
type MutltipleChoiceQuestion struct {
	QuestionID   int
	Label        string
	QuestionType string
}

// TrueOrFalseQuestion ...
type TrueOrFalseQuestion struct {
	QuestionID   int
	Label        string
	QuestionType string
	AnswerIs     bool
}

// SliderQuestion ...
type SliderQuestion struct {
	QuestionID   int
	Label        string
	QuestionType string
	AnswerMin    int
	AnswerMax    int
}

// MutltipleChoiceQuestionAnswer ...
type MutltipleChoiceQuestionAnswer struct {
	ID                  int
	MutltipleQuestionID int
	Label               string
	IsCorrect           bool
}

// Quiz ...
type Quiz struct {
	ID     int
	Label  string
	Author int
}

// QuizQuestion ...
type QuizQuestion struct {
	QuizID     int
	QuestionID int
}
