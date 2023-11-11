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

func AddQuestions(c *gin.Context) {
	var question domain.Questions
	if err := c.ShouldBindJSON(&question); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(question)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	questionResponse, err := usecase.AddQuestions(question)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the question", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added question", questionResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func UpdateQuestion(c *gin.Context) {
	var p models.SetNewName

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	a, err := usecase.UpdateQuestion(p.Current, p.New)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the Question", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully renamed the Question", a, nil)
	c.JSON(http.StatusOK, successRes)

}

func DeleteQuestion(c *gin.Context) {
	questionId := c.Query("id")
	err := usecase.DeleteQuestion(questionId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the question", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
