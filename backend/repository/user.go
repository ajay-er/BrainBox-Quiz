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

	var QuizNames models.QuizesInCategopry
	if err := database.DB.Raw("select quiz_name from quizes where category_id = ? limit ? offset ?", id, count, offset).Scan(&QuizNames.QuizName).Error; err != nil {
		return models.QuizesInCategopry{}, err
	}
	return QuizNames, nil
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

func GetOptionsFromQuestionIds(questionIds []uint) ([]models.OptionsResponse, error) {
	var options []models.OptionsResponse
	if err := database.DB.Raw("select * from options where question_id IN ?", questionIds).Scan(&options).Error; err != nil {
		return []models.OptionsResponse{}, err
	}
	return options, nil
}

func GetOptionsFromQuestionId(questionIds []uint) ([]models.OptionsResponse, error) {
	var options []models.OptionsResponse

	if err := database.DB.Where("question_id IN (?)", questionIds).Find(&options).Error; err != nil {
		return nil, err
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


