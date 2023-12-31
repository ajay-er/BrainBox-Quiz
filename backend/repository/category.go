package repository

import (
	database "backend/db"
	"backend/domain"
	"errors"
	"strconv"
)

func AddCategory(category domain.Category) (domain.Category, error) {
	var b string
	var categoryResponse domain.Category
	if err := database.DB.Raw("INSERT INTO categories (category_name, icon_svg) VALUES (?, ?) RETURNING id, category_name, icon_svg", category.CategoryName, category.IconSvg).Scan(&categoryResponse).Error; err != nil {
		return domain.Category{}, err
	}
	err := database.DB.Raw("SELECT id ,category_name,icon_svg FROM categories WHERE category_name = ?", b).Scan(&categoryResponse).Error
	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil
}
func CheckCategory(current string) (bool, error) {
	var i int
	err := database.DB.Raw("select count(*) from categories where category_name =?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	}
	return true, err

}
func UpdateCategory(current, new string) (domain.Category, error) {
	if database.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}
	if err := database.DB.Exec("update categories set category_name =$1 where category_name =$2", new, current).Error; err != nil {
		return domain.Category{}, err
	}
	var newcat domain.Category
	if err := database.DB.First(&newcat, "category_name=?", new).Error; err != nil {
		return domain.Category{}, err
	}
	return newcat, nil

}
func DeleteCategory(categoryID string) error {
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return errors.New("couldn't convert")
	}
	result := database.DB.Exec("delete from categories where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
