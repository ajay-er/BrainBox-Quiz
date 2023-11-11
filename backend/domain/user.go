package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model `json:"-"`
	ID          uint   `json:"id" gorm:"unique;not null"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email" validate:"email"`
	Password    string `json:"password" validate:"min=8,max=20"`
	Phone       string `json:"phone"`
	Blocked     bool   `json:"blocked" gorm:"default:false"`
	Isadmin     bool   `json:"is_admin" gorm:"default:false"`
}

type Category struct {
	*gorm.Model  `json:"-"`
	ID           uint       `json:"id" gorm:"unique; not null"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index"`
	CategoryName string     `json:"category_name" gorm:"unique; "  `
}

type Quizes struct {
	*gorm.Model `json:"-"`
	ID          uint     `json:"id" gorm:"unique; not null"`
	QuizName    string   `json:"quiz_name" gorm:"unique;" `
	Description string `json:"description"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryId;references:ID;constraint:OnDelete:CASCADE"`
}

type Questions struct {
	*gorm.Model `json:"-"`
	ID          uint   `json:"id" gorm:"unique; not null"`
	Question    string `json:"question" `
	QuizId      uint   `json:"quiz_id"`
	Quizes      Quizes `json:"-" gorm:"foreignKey:QuizId;references:ID;constraint:OnDelete:CASCADE"`
}

type Options struct {
	*gorm.Model `json:"-"`
	ID          uint      `json:"id" gorm:"unique; not null"`
	Option      string    `json:"option" `
	QuestionId  uint      `json:"question_id"`
	Questions   Questions `json:"-" gorm:"foreignKey:QuestionId;references:ID;constraint:OnDelete:CASCADE"`
	IsCorrect   bool      `json:"is_correct" default:"false"`
}

type QuizResults struct {
	*gorm.Model `json:"-"`
	ID          uint   `json:"id" gorm:"unique; not null"`
	Score       int    `json:"score"`
	QuizId      uint   `json:"quiz_id"`
	Quizes      Quizes `json:"-" gorm:"foreignKey:QuizId;references:ID;constraint:OnDelete:CASCADE"`
	UserId      uint   `json:"user_id"`
	User        User   `json:"-" gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
}
type CreateQuiz struct {
	CategoryName string
	QuizName     string
	Question     []QuizQuestion
}
type QuizQuestion struct {
	Questions string
	Options   []OptionValues
}
type OptionValues struct {
	Option    string
	IsCorrect bool
}
