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
