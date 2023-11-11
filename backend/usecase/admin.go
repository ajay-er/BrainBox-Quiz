package usecase

import (
	"backend/domain"
	"backend/helpers"
	"backend/models"
	"backend/repository"
	"errors"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(adminDetail models.AdminLogin) (domain.TokenAdmin, error) {

	adminCompareDetails, err := repository.AdminLogin(adminDetail)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetail.Password))
	if err != nil {
		return domain.TokenAdmin{}, errors.New("password not matching")
	}
	var adminDetailResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("password not matching")
	}
	tokenString, err := helpers.GenerateTokenAdmin(adminDetailResponse)

	if err != nil {
		return domain.TokenAdmin{}, errors.New("error in creating token")
	}
	return domain.TokenAdmin{
		Admin: adminDetailResponse,
		Token: tokenString,
	}, nil

}
func AdminDashboard() (models.TotalAdminDashboard, error) {

	userdetails, err := repository.DashboardUserDetails()
	if err != nil {
		return models.TotalAdminDashboard{}, err
	}
	categorydetails, err := repository.DashBoardCategoryDetails()
	if err != nil {
		return models.TotalAdminDashboard{}, err
	}
	// quizdetails,err:=repository.DashboardquizDetails()
	// if err!=nil{
	// 	return models.TotalAdminDashboard{},err
	// }
	// questiondetails,err:=repository.DashboardquestionDetails()
	// if err!=nil{
	// 	return models.TotalAdminDashboard{},err
	// }
	return models.TotalAdminDashboard{
		DashBoardUser:     userdetails,
		DashBoardCategory: categorydetails,
	}, nil
}

