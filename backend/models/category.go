package models

type Category struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"`
}

type CategoryDetails struct {
	Categories []Category
}

type QuizesInCategopry struct {
	QuizName []string

	TotalQuizes int
}

type QuizResponse struct {
	ID         uint   `json:"id" `
	QuizName   string `json:"quiz_name"`
	CategoryId uint   `json:"category_id"`
}
type QuestionsResponse struct {
	ID       uint   `json:"id" `
	Question string `json:"question"`
	QuizId   uint   `json:"quiz_id"`
}

type OptionsResponse struct {
	ID         uint   `json:"id"`
	Option     string `json:"option"`
	QuestionId uint   `json:"question_id"`
}

type TotalQuizResponse struct{
	QuizResponse QuizResponse
	QuestionsResponse []QuestionsResponse
	OptionsResponse []OptionsResponse
}