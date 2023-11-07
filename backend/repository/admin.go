package repository

import (
	database "backend/db"
	"backend/domain"
	"backend/models"
	"errors"
)

func AdminLogin(adminDetail models.AdminLogin) (domain.Admin, error) {
	var count int
	err := database.DB.Raw("select count(*) from users where email = ? and isadmin=true", adminDetail.Email).Scan(&count).Error
	if err != nil {
		return domain.Admin{}, err
	}

	if count > 0 {

		var adminCompareDetail domain.Admin

		if err := database.DB.Raw("select * from users where email = ? and isadmin=true", adminDetail.Email).Scan(&adminCompareDetail).Error; err != nil {
			return domain.Admin{}, err
		}
		return adminCompareDetail, nil
	}
	return domain.Admin{}, errors.New("the admin login credential is not valid")
}
func DashboardUserDetails() (models.DashBoardUser, error) {
	var userDetail models.DashBoardUser
	err := database.DB.Raw("select count(*) from users where isadmin=false").Scan(&userDetail.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = database.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetail.BlockedUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	var userDetails []models.SignupDetailResponse
	if err := database.DB.Raw("select * from users where isadmin=false").Scan(&userDetails).Error; err != nil {
		return models.DashBoardUser{}, err
	}
	return models.DashBoardUser{
		TotalUsers:   userDetail.TotalUsers,
		BlockedUsers: userDetail.BlockedUsers,
		Users:        userDetails,
	}, nil

}

func DashBoardCategoryDetails() (models.DashBoardCategory,error){
var categorydetail models.DashBoardCategory
err := database.DB.Raw("select count(*) from categories").Scan(&categorydetail.TotalCategories).Error
	if err != nil {
		return models.DashBoardCategory{}, nil
	}
	var categoryDetails []models.Category
	if err := database.DB.Raw("select * from categories").Scan(&categoryDetails).Error; err != nil {
		return models.DashBoardCategory{}, err
	}
	return models.DashBoardCategory{
		TotalCategories: categorydetail.TotalCategories,
		Category: categoryDetails,
	},nil
}
