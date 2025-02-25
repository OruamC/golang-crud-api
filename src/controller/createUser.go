package controller

import (
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/configuration/validation"
	"github.com/OruamC/golang-crud-api/src/controller/model/request"
	"github.com/OruamC/golang-crud-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))
	responseUser := response.UserResponse{
		ID:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusCreated, responseUser)
}
