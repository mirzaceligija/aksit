package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mirzaceligija/askit/dto"
	"github.com/mirzaceligija/askit/entity"
	"github.com/mirzaceligija/askit/helper"
	"github.com/mirzaceligija/askit/service"
)

type QuestionController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindByUserID(context *gin.Context)
}

type questionController struct {
	questionService service.QuestionService
	jwtService      service.JWTService
}

func NewQuestionController(questionServ service.QuestionService, jwtServ service.JWTService) QuestionController {
	return &questionController{
		questionService: questionServ,
		jwtService:      jwtServ,
	}
}

func (c *questionController) All(context *gin.Context) {
	orderBy := context.Query("orderBy")
	page, err := strconv.ParseUint(context.Query("page"), 0, 0)
	if err != nil || orderBy == "" {
		res := helper.BuildErrorResponse("No param was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var questions []entity.Question = c.questionService.All(page, orderBy)
	res := helper.BuildResponse(true, "OK", questions)
	context.JSON(http.StatusOK, res)
}

func (c *questionController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var question entity.Question = c.questionService.FindByID(id)
	if (question == entity.Question{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", question)
		context.JSON(http.StatusOK, res)
	}
}

func (c *questionController) Insert(context *gin.Context) {
	var questionCreateDTO dto.QuestionCreateDTO
	errDTO := context.ShouldBind(&questionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			questionCreateDTO.UserID = convertedUserID
		}
		result := c.questionService.Insert(questionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *questionController) Update(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	var questionUpdateDTO dto.QuestionUpdateDTO
	questionUpdateDTO.ID = id
	errDTO := context.ShouldBind(&questionUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.questionService.IsAllowedToEdit(userID, questionUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			questionUpdateDTO.UserID = id
		}
		result := c.questionService.Update(questionUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *questionController) Delete(context *gin.Context) {
	var question entity.Question
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	question.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.questionService.IsAllowedToEdit(userID, question.ID) {
		c.questionService.Delete(question)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *questionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *questionController) FindByUserID(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	id, err := strconv.ParseUint(userID, 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var questions []entity.Question = c.questionService.FindByUserID(id)
	res := helper.BuildResponse(true, "OK", questions)
	context.JSON(http.StatusOK, res)
}
