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

func GetUsers(c *gin.Context) {

	pageStr := c.Param("page")

	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	countStr := c.Query("count")

	if countStr == "" {
		countStr = "0"
	}

	pageSize, err := strconv.Atoi(countStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := usecase.GetUsers(page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}
