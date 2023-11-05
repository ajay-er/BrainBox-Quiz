package usecase

import (
	"backend/helpers"
	"backend/models"
	"backend/repository"
	"errors"
	"fmt"
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
