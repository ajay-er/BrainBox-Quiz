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

