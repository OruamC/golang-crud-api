package controller

import (
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/configuration/validation"
	"github.com/OruamC/golang-crud-api/src/controller/model/request"
	"github.com/OruamC/golang-crud-api/src/model"
	"github.com/OruamC/golang-crud-api/src/model/service"
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))
	c.String(http.StatusCreated, "")
}
