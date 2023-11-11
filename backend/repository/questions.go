package repository

import (
	database "backend/db"
	"backend/domain"
	"errors"
	"fmt"
	"strconv"
)

func AddQuestions(question domain.Questions) (domain.Questions, error) {

	var b string
	query := "INSERT INTO questions (question, quiz_id) VALUES (?, ?) RETURNING question"

	err := database.DB.Raw(query, question.Question, question.QuizId).Scan(&b).Error
	fmt.Println(b, "ghjk")

	if err != nil {
		return domain.Questions{}, err
	}
	var questionResponse domain.Questions
	err = database.DB.Raw("SELECT id ,question,quiz_id FROM questions WHERE question = ?", b).Scan(&questionResponse).Error
	if err != nil {
		return domain.Questions{}, err
	}

	return questionResponse, nil
}

func CheckQuestion(current string) (bool, error) {
	var i int
	err := database.DB.Raw("select count(*) from questions where question =?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	}
	return true, err

}

func UpdateQuestion(current, new string) (domain.Questions, error) {
	if database.DB == nil {
		return domain.Questions{}, errors.New("database connection is nil")
	}
	if err := database.DB.Exec("update questions set question =$1 where question =$2", new, current).Error; err != nil {
		return domain.Questions{}, err
	}
	var newcat domain.Questions
	if err := database.DB.First(&newcat, "question=?", new).Error; err != nil {
		return domain.Questions{}, err
	}
	return newcat, nil

}

func DeleteQuestion(questionId string) error {
	id, err := strconv.Atoi(questionId)
	if err != nil {
		return errors.New("couldn't convert")
	}
	result := database.DB.Exec("delete from questions where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
