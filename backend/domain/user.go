package domain

import "gorm.io/gorm"

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
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"`
}

type Quizes struct {
	*gorm.Model `json:"-"`
	ID          uint     `json:"id" gorm:"unique; not null"`
	QuizName    string   `json:"quiz_name"`
	CaregoryId  uint     `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignKey:CaregoryId;references:ID;constraint:OnDelete:CASCADE"`
}

type Questions struct {
	*gorm.Model `json:"-"`
	ID          uint   `json:"id" gorm:"unique; not null"`
	Question    string `json:"question"`
	QuizId      uint   `json:"quiz_id"`
	Quizes      Quizes `json:"-" gorm:"foreignKey:QuizId;references:ID;constraint:OnDelete:CASCADE"`
}

type Options struct {
	*gorm.Model `json:"-"`
	ID          uint      `json:"id" gorm:"unique; not null"`
	Opton       string    `json:"option"`
	QuestionId  uint      `json:"question_id"`
	Questions   Questions `json:"-" gorm:"foreignKey:QuestionId;references:ID;constraint:OnDelete:CASCADE"`
}
