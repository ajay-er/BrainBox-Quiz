package handlers

import (
	"backend/models"
	"backend/response"
	"backend/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AdminLogin(c *gin.Context) {
	var adminDetail models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetail); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := usecase.AdminLogin(adminDetail)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}
func AdminDashboard(c *gin.Context) {
	dashboard, err := usecase.AdminDashboard()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "dashboard could not be displayed", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "dashboard displayed succesfully", dashboard, nil)
	c.JSON(200, successRes)
	return
}
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user id is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}

	var user models.UsersProfileDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err = validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println(user, "✔")
	updatedDetails, err := usecase.UpdateUserDetails(user, id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed update user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Updated User Details", updatedDetails, nil)
	c.JSON(http.StatusOK, successRes)
}
func DeleteUser(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = usecase.DeleteUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully deleted the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func BlockUser(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	err = usecase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func UnBlockUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	err = usecase.UnBlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldn't blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}