package handlers

import (
	"backend/models"
	"backend/response"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {
	var userSignup models.SignupDetail
	err := c.ShouldBindJSON(&userSignup)
	if err != nil {
		errRes := response.ClientResponse(400, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//creating a newuser signup with the given deatil passing into the bussiness logic layer
	userCreated, err := usecase.UserSignup(userSignup)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong formaaaaat", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)
}
func UserLoginWithPassword(c *gin.Context) {
	var userLoginDetail models.LoginDetail
	if err := c.ShouldBindJSON(&userLoginDetail); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := validator.New().Struct(userLoginDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userLoggedInWithPassword, err := usecase.UserLoginWithPassword(userLoginDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully Logged In With password", userLoggedInWithPassword, nil)
	c.JSON(http.StatusCreated, successRes)

}

func Home(c *gin.Context) {

	getCategoryDetails, err := usecase.GetCategory()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in taking category details from db", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	successRes := response.ClientResponse(http.StatusCreated, "User successfully Logged In With password", getCategoryDetails, nil)
	c.JSON(http.StatusCreated, successRes)

}

func Categories(c *gin.Context) {
	category_id := c.Query("id")
	page_no := c.Query("page")
	count_no := c.Query("count")
	page, err := strconv.Atoi(page_no)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion of page", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	count, err := strconv.Atoi(count_no)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion of count", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	quizNames, err := usecase.Categories(category_id, page, count)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in taking quizes from db", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "Successfully get the quiz details", quizNames, nil)
	c.JSON(http.StatusCreated, successRes)

}


func Quizes(c *gin.Context) {

	quizId := c.Query("quiz_id")
	totalResponse, err := usecase.Quizes(quizId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in getting qns and options", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	successRes := response.ClientResponse(http.StatusCreated, "Successfully get the quiz,qns and options", totalResponse, nil)
	c.JSON(http.StatusCreated, successRes)
}

