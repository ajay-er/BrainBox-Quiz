package handlers

import (
	"backend/models"
	"backend/response"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
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
