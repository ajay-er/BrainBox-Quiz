package handlers

import (
	"backend/models"
	"backend/response"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
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
