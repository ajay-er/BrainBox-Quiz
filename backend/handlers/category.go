package handlers

import (
	"backend/domain"
	"backend/models"
	"backend/response"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	categoryResponse, err := usecase.AddCategory(category)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Category", categoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}
func UpdateCategory(c *gin.Context) {
	var p models.SetNewName

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	a, err := usecase.UpdateCategory(p.Current, p.New)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully renamed the category", a, nil)
	c.JSON(http.StatusOK, successRes)

}
func DeleteCategory(c *gin.Context) {
	categoryID := c.Query("id")
	err := usecase.DeleteCategory(categoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the Category", nil, nil)
	c.JSON(http.StatusOK, successRes)
}