package handlers

import (
	"backend/domain"
	"backend/models"
	"backend/response"
	"backend/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



func CreateQuiz(c *gin.Context) {

	var quiz domain.CreateQuiz

	if err := c.ShouldBindJSON(&quiz); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println("quizis ", quiz)
	err := validator.New().Struct(quiz)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = usecase.CreateQuiz(quiz)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the quiz", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added quiz", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func UpdateQuiz(c *gin.Context) {
	var p models.SetNewName

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	a, err := usecase.UpdateQuiz(p.Current, p.New)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the quiz", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully renamed the quiz", a, nil)
	c.JSON(http.StatusOK, successRes)

}

func DeleteQuiz(c *gin.Context) {
	quizId := c.Query("id")
	err := usecase.DeleteQuiz(quizId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the quiz", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
