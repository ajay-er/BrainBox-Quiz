package repository

import (
	database "backend/db"
	"backend/domain"
	"backend/models"
	"errors"
	"strconv"
)

func AddCategoryByName(CategoryName string) (models.CategoryResponse, error) {
	var categoryResponse models.CategoryResponse
	err := database.DB.Raw("select * from categories where category_name=?", CategoryName).Scan(&categoryResponse).Error
	if err != nil {
		return models.CategoryResponse{}, err
	}
	return categoryResponse, nil
}
func AddQuizByName(quizname string, categoryId uint) (models.QuizResponse, error) {
	var quizResponse models.QuizResponse
	query := "INSERT INTO quizes (quiz_name, category_id) VALUES (?, ?) RETURNING id ,quiz_name,category_id"

	err := database.DB.Raw(query, quizname, categoryId).Scan(&quizResponse).Error

	if err != nil {
		return models.QuizResponse{}, err
	}

	return quizResponse, nil
}
func AddQuestionsByQuestionName(Questions string, quizId uint) (models.QuestionsResponse, error) {

	var questionResponse models.QuestionsResponse
	query := "INSERT INTO questions (question, quiz_id) VALUES (?, ?) RETURNING id ,question,quiz_id"

	err := database.DB.Raw(query, Questions, quizId).Scan(&questionResponse).Error

	if err != nil {
		return models.QuestionsResponse{}, err
	}

	return questionResponse, nil

}
func AddOptionsByOptionName(Option string, IsCorrect bool, ID uint) (models.OptionsResponse, error) {
	var optionsRespo models.OptionsResponse
	query := "insert into options(option,question_id,is_correct) values(?,?,?) RETURNING id ,option,question_id,is_correct"
	err := database.DB.Raw(query, Option, ID, IsCorrect).Scan(&optionsRespo).Error
	if err != nil {
		return models.OptionsResponse{}, err
	}
	return optionsRespo, nil
}
func CheckQuiz(current string) (bool, error) {
	var i int
	err := database.DB.Raw("select count(*) from quizes where quiz_name =?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	}
	return true, err

}

func UpdateQuiz(current, new string) (domain.Quizes, error) {
	if database.DB == nil {
		return domain.Quizes{}, errors.New("database connection is nil")
	}
	if err := database.DB.Exec("update quizes set quiz_name =$1 where quiz_name =$2", new, current).Error; err != nil {
		return domain.Quizes{}, err
	}
	var newcat domain.Quizes
	if err := database.DB.First(&newcat, "quiz_name=?", new).Error; err != nil {
		return domain.Quizes{}, err
	}
	return newcat, nil

}

func DeleteQuiz(quizId string) error {
	id, err := strconv.Atoi(quizId)
	if err != nil {
		return errors.New("couldn't convert")
	}
	result := database.DB.Exec("delete from quizes where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
