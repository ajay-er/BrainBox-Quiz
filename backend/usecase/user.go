package usecase

import (
	"backend/helpers"
	"backend/models"
	"backend/repository"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(user models.SignupDetail) (*models.TokenUser, error) {

	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email, "🙌")
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := repository.CheckUserExistsByPhone(user.Phone)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}

	//if the signing up is a new user then hashing the password
	hashedPassword, err := helpers.PasswordHashing(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}

	user.Password = hashedPassword

	//after hashing adding the user detail into the database and taking the added user detail to the userdata
	userData, err := repository.UserSignup(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add the user ")
	}

	//creating a jwt token for the new user with the detail that has been stored in the database
	accessToken, err := helpers.GenerateAccessToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create access token due to error")
	}
	refreshToken, err := helpers.GenerateRefreshToken(userData)
	if err != nil {

		return &models.TokenUser{}, errors.New("counldnt create refresh token due to error")

	}

	return &models.TokenUser{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func UserLoginWithPassword(user models.LoginDetail) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return &models.TokenUser{}, errors.New("email  does not exsist")
	}

	userDetails, err := repository.FindUserDetailsByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))

	if err != nil {
		return &models.TokenUser{}, errors.New("password not matching")
	}
	var user_details models.SignupDetailResponse
	err = copier.Copy(&user_details, &userDetails)
	if err != nil {
		return &models.TokenUser{}, err
	}
	accessToken, err := helpers.GenerateAccessToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not create accesstoken due to internal error")
	}
	refreshToken, err := helpers.GenerateRefreshToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create refresh token due to error")
	}

	return &models.TokenUser{
		Users:        user_details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}