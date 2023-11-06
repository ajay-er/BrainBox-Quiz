package usecase

import (
	"backend/domain"
	"backend/helpers"
	"backend/models"
	"backend/repository"

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
		return domain.TokenAdmin{}, err
	}
	var adminDetailResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	tokenString, err := helpers.GenerateTokenAdmin(adminDetailResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}
	return domain.TokenAdmin{
		Admin: adminDetailResponse,
		Token: tokenString,
	}, nil

}
