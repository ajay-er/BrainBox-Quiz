package usecase

import (
	"backend/domain"
	"backend/repository"
	"errors"
)

func AddQuestions(question domain.Questions) (domain.Questions, error) {
	questionResponse, err := repository.AddQuestions(question)
	if err != nil {
		return domain.Questions{}, err
	}
	return questionResponse, nil
}
func UpdateQuestion(current string, new string) (domain.Questions, error) {
	result, err := repository.CheckQuestion(current)
	if err != nil {
		return domain.Questions{}, err
	}
	if !result {
		return domain.Questions{}, errors.New("there is no question name as you mentioned")
	}
	newCat, err := repository.UpdateQuestion(current, new)
	if err != nil {
		return domain.Questions{}, err
	}
	return newCat, err
}

func DeleteQuestion(questionId string) error {
	err := repository.DeleteQuestion(questionId)
	if err != nil {
		return err
	}
	return nil
}
