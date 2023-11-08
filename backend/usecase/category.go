package usecase

import (
	"backend/domain"
	"backend/repository"
	"errors"
)

func AddCategory(category domain.Category) (domain.Category, error) {
	productResponse, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return productResponse, nil
}
func UpdateCategory(current string, new string) (domain.Category, error) {
	result, err := repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}
	if !result {
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}
	newCat, err := repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCat, err
}
func DeleteCategory(categoryID string) error {
	err := repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}