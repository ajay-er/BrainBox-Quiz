package usecase

import (
	"backend/domain"
	"backend/repository"
	"errors"
	"fmt"
)

func CreateQuiz(quiz domain.CreateQuiz) error {

	categoryResponse, err := repository.AddCategoryByName(quiz.CategoryName)
	if err != nil {
		return err
	}

	quizResponse, err := repository.AddQuizByName(quiz.QuizName, categoryResponse.ID)
	if err != nil {
		return err
	}

	fmt.Println(quiz.Question, "hiii")
	for _, question := range quiz.Question {

		questionResponse, err := repository.AddQuestionsByQuestionName(question.Questions, quizResponse.ID)
		fmt.Println(questionResponse, "😜")

		if err != nil {
			return err
		}

		for _, option := range question.Options {

			fmt.Println(option.Option, "😂", option.IsCorrect, "wda")
			OptionsResponse, err := repository.AddOptionsByOptionName(option.Option, option.IsCorrect, questionResponse.ID)
			fmt.Println(OptionsResponse)
			if err != nil {

				return err
			}
		}

	}
	fmt.Println("hhhhhh")
	return nil
}
func UpdateQuiz(current string, new string) (domain.Quizes, error) {
	result, err := repository.CheckQuiz(current)
	if err != nil {
		return domain.Quizes{}, err
	}
	if !result {
		return domain.Quizes{}, errors.New("there is no quiz name as you mentioned")
	}
	newCat, err := repository.UpdateQuiz(current, new)
	if err != nil {
		return domain.Quizes{}, err
	}
	return newCat, err
}

func DeleteQuiz(quizId string) error {
	err := repository.DeleteQuiz(quizId)
	if err != nil {
		return err
	}
	return nil
}
