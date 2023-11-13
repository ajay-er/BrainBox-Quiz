package models

type Category struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"`
	IconSvg      string `json:"icon_svg"`
	TotalQuizes int `json:"total_quizes"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type CategoryDetails struct {
	Categories []Category
}

type QuizesInCategopry struct {
	Quizdetail  []QuizResponse
	TotalQuizes int
}
type CategoryResponse struct {
	ID           uint   `json:"id" `
	CategoryName string `json:"category_name"`
}
type QuizResponse struct {
	ID          uint   `json:"id" `
	QuizName    string `json:"quiz_name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
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
	IsCorrect  bool   `json:"is_correct"`
}

// type TotalQuizResponse struct {
// 	QuizResponse      QuizResponse
// 	QuestionsResponse []QuestionsResponse
// 	OptionsResponse   []OptionsResponse
// }

// type QuestionAndOption struct {
// 	QuestionId int
// 	Question   string
// 	Options    []OptionsResponse
// }

type QuizQuestion struct {
	Question QuestionsResponse
	Options  []OptionsResponse
}
type TotalQuizResponse struct {
	Category  CategoryResponse
	Quiz      QuizResponse
	Questions []QuizQuestion
}
type ScoreResponse struct {
	Score       int     `json:"score"`
	Percentage  float64 `json:"percentage"`
	Description string  `json:"descrtiption"`
}
