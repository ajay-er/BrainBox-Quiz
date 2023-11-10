package repository

import (
	database "backend/db"
	"backend/domain"
	"backend/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := database.DB.Where(&domain.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
func CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func UserSignup(user models.SignupDetail) (models.SignupDetailResponse, error) {
	var signupDetail models.SignupDetailResponse
	err := database.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,phone", user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Scan(&signupDetail).Error
	if err != nil {
		fmt.Println("Repository error:", err)
		return models.SignupDetailResponse{}, err
	}
	return signupDetail, nil

}
func FindUserDetailsByEmail(user models.LoginDetail) (models.UserLoginResponse, error) {
	var userdetails models.UserLoginResponse

	err := database.DB.Raw(
		`SELECT * FROM users where email = ? and blocked = false`, user.Email).Scan(&userdetails).Error

	if err != nil {
		return models.UserLoginResponse{}, errors.New("error checking user details")
	}
	return userdetails, nil

}

func CheckUserAvailability(email string) bool {
	fmt.Println(email, "🤣")
	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := database.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	fmt.Println(count, "🎶")
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func UpdateUserEmail(email string, userID int) error {

	err := database.DB.Exec("update users set email = ? where id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateUserPhone(phone string, userID int) error {

	err := database.DB.Exec("update users set phone = ? where id = ?", phone, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateFirstName(name string, userID int) error {

	err := database.DB.Exec("update users set firstname = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}
func UpdateLastName(name string, userID int) error {

	err := database.DB.Exec("update users set lastname = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}
func UserDetails(userID int) (models.UsersProfileDetails, error) {

	var userDetails models.UsersProfileDetails
	err := database.DB.Raw("select users.firstname,users.lastname,users.email,users.phone from users  where users.id = ?", userID).Row().Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Email, &userDetails.Phone)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}
	return userDetails, nil
}
func GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	var userDetails []models.UserDetailsAtAdmin

	if page <= 0 {
		page = 1
	}

	if count <= 0 {
		count = 6
	}
	offset := (page - 1) * count

	if err := database.DB.Raw("select id,firstname,lastname,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {

		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}
func DeleteUser(id int) error {
	result := database.DB.Exec("delete from users where id = ?", id)
	if result.RowsAffected < 1 {
		return errors.New("no records with that is exist")
	}
	return nil
}
func GetUserByID(id int) (*domain.User, error) {

	var user domain.User
	result := database.DB.Where(&domain.User{ID: uint(id)}).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	if err := database.DB.Raw("select * from users where id=?", id).Scan(&user).Error; err != nil {
		return nil, nil
	}
	return &user, nil
}
func UpdateBlockUserByID(user *domain.User) error {
	err := database.DB.Exec("update users set blocked=? where id=?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategory() (models.CategoryDetails, error) {
	var categoryDetails models.CategoryDetails
	if err := database.DB.Raw("select * from categories").Scan(&categoryDetails.Categories).Error; err != nil {

		return models.CategoryDetails{}, err
	}
	return categoryDetails, nil
}

func GetAllQuizesByCategoryID(id string, page int, count int) (models.QuizesInCategopry, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var quiz []models.QuizResponse
	if err := database.DB.Raw("select * from quizes where category_id = ? limit ? offset ?", id, count, offset).Scan(&quiz).Error; err != nil {
		return models.QuizesInCategopry{}, err
	}
	return models.QuizesInCategopry{
		Quizdetail: quiz,
	}, nil
}

func GetTotalNumberOfQuizInACategory(id string) (int, error) {
	var totalquizes int
	if err := database.DB.Raw("select count(*) from quizes where category_id = ? ", id).Scan(&totalquizes).Error; err != nil {
		return 0, err
	}
	return totalquizes, nil

}

func GetQuizDetailsFromQuizId(id string) (models.QuizResponse, error) {
	var quizResponse models.QuizResponse
	if err := database.DB.Raw("select * from quizes where id = ?", id).Scan(&quizResponse).Error; err != nil {
		return models.QuizResponse{}, err
	}
	return quizResponse, nil
}
func GetMatchingQuestionsFromQuizId(id string) ([]models.QuestionsResponse, error) {
	var questionDetails []models.QuestionsResponse
	if err := database.DB.Raw("select * from questions where quiz_id = ?", id).Scan(&questionDetails).Error; err != nil {
		return []models.QuestionsResponse{}, err
	}
	return questionDetails, nil

}

// func GetQuestionDetailsFromQuestionId(id string) (models.QuestionsResponse, error) {
// 	var questionDetails models.QuestionsResponse
// 	if err := database.DB.Raw("select * from questions where question_id = ?", id).Scan(&questionDetails).Error; err != nil {
// 		return models.QuestionsResponse{}, err
// 	}
// 	return questionDetails, nil

// }

func GetOptionsFromQuestionIds(questionIds uint) ([]models.OptionsResponse, error) {
	var options []models.OptionsResponse
	if err := database.DB.Raw("select * from options where question_id = ?", questionIds).Scan(&options).Error; err != nil {
		return []models.OptionsResponse{}, err
	}
	return options, nil
}

func GetOptionById(optionId string) (models.OptionsResponse, error) {

	var options models.OptionsResponse
	if err := database.DB.Raw("select * from options where id = ?", optionId).Scan(&options).Error; err != nil {
		return models.OptionsResponse{}, err
	}
	return options, nil
}
func UpdateScore(userId string, score int, quizId string) error {
	if err := database.DB.Exec("insert into quiz_results(user_id,score,quiz_id) values(?,?,?)", userId, score, quizId).Error; err != nil {
		return err
	}
	return nil
}
