package usecase

import (
	"backend/helpers"
	"backend/models"
	"backend/repository"
	"errors"
	"fmt"
	"regexp"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func IsEmailValid(email string) bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	if match {
		return true
	} else {
		return false
	}
}
func IsValidPhoneNumber(phoneNumber string) bool {

	phoneRegex := `^[789]\d{9}$`
	match, _ := regexp.MatchString(phoneRegex, phoneNumber)
	if match {
		return true
	} else {
		return false
	}
}
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
func GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := repository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}
func UpdateUserDetails(user models.UsersProfileDetails, user_id int) (models.UsersProfileDetails, error) {
	// if !IsEmailValid(user.Email) {
	// 	return models.UsersProfileDetails{}, errors.New("invalid email format")
	// }

	// if !IsValidPhoneNumber(user.Phone) {
	// 	return models.UsersProfileDetails{}, errors.New("invalid phone number format")
	// }
	userExist := repository.CheckUserAvailability(user.Email)
	// update with email that does not already exist
	if userExist {
		return models.UsersProfileDetails{}, errors.New("user already exist, choose different email")
	}
	userExistByPhone, err := repository.CheckUserExistsByPhone(user.Phone)

	if err != nil {
		return models.UsersProfileDetails{}, errors.New("error with server")
	}
	if userExistByPhone != nil {
		return models.UsersProfileDetails{}, errors.New("user with this phone is already exists")
	}
	// which all field are not empty (which are provided from the front end should be updated)
	if user.Email != "" {
		repository.UpdateUserEmail(user.Email, user_id)
	}

	if user.Firstname != "" {
		repository.UpdateFirstName(user.Firstname, user_id)
	}
	if user.Lastname != "" {
		repository.UpdateLastName(user.Lastname, user_id)
	}

	if user.Phone != "" {
		repository.UpdateUserPhone(user.Phone, user_id)
	}

	return repository.UserDetails(user_id)
}
func DeleteUser(id int) error {
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func BlockUser(id int) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}
func UnBlockUser(id int) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if !user.Blocked {
		return errors.New("user is already unblocked")
	} else {
		user.Blocked = true
	}
	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory() (models.CategoryDetails, error) {
	var categoryDetails models.CategoryDetails

	categoryDetails, err := repository.GetCategory()
	if err != nil {
		return models.CategoryDetails{}, err
	}
	return categoryDetails, nil

}

func Categories(id string, page int, count int) (models.QuizesInCategopry, error) {

	quizNames, err := repository.GetAllQuizesByCategoryID(id, page, count)
	if err != nil {
		return models.QuizesInCategopry{}, err
	}
	totalQuizes, err := repository.GetTotalNumberOfQuizInACategory(id)
	if err != nil {
		return models.QuizesInCategopry{}, err
	}
	quizNames.TotalQuizes = totalQuizes

	return quizNames, nil

}

func Quizes(id string) (models.TotalQuizResponse, error) {

	quizResponse, err := repository.GetQuizDetailsFromQuizId(id)
	if err != nil {
		return models.TotalQuizResponse{}, err
	}
	questionDetails, err := repository.GetMatchingQuestionsFromQuizId(id)
	if err != nil {
		return models.TotalQuizResponse{}, err
	}

	totalResponse := models.TotalQuizResponse{
		Quiz:      quizResponse,
		Questions: []models.QuizQuestion{},
	}
	for _, question := range questionDetails {

		optionsResponse, err := repository.GetOptionsFromQuestionIds(question.ID)
		if err != nil {
			return models.TotalQuizResponse{}, err
		}
		questionWithOptions := models.QuizQuestion{
			Question: question,
			Options:  optionsResponse,
		}

		totalResponse.Questions = append(totalResponse.Questions, questionWithOptions)

	}

	return totalResponse, nil

}
func ScoreTracking(optionId []string) (models.ScoreResponse, error) {
	var count = 0

	for _, option := range optionId {
		options, err := repository.GetOptionById(option)
		if err != nil {
			return models.ScoreResponse{}, err
		}
		if options.IsCorrect {
			count++
		}
	}
	var scoreResponse models.ScoreResponse
	scoreResponse.Score = count
	scoreResponse.Percentage = (float64(scoreResponse.Score) / float64(len(optionId)) * 100)
	fmt.Println(scoreResponse.Percentage)
	fmt.Println(scoreResponse.Score)
	fmt.Println(len(optionId))

	if scoreResponse.Percentage == 100 {
		scoreResponse.Description = "Excellent"
	}
	if scoreResponse.Percentage < 100 && scoreResponse.Percentage >= 90 {
		scoreResponse.Description = "Very Good"
	}
	if scoreResponse.Percentage < 90 && scoreResponse.Percentage >= 80 {
		scoreResponse.Description = " Good"
	}
	if scoreResponse.Percentage < 80 && scoreResponse.Percentage >= 50 {
		scoreResponse.Description = "Fine"
	}
	if scoreResponse.Percentage < 50 {
		scoreResponse.Description = "Need to impr"
	}

	return scoreResponse, nil

}
