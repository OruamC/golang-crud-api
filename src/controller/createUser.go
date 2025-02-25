package controller

import (
	"fmt"
	"github.com/OruamC/golang-crud-api/src/configuration/validation"
	"github.com/OruamC/golang-crud-api/src/controller/model/request"
	"github.com/OruamC/golang-crud-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	log.Println("Init createUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error: %s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	responseUser := response.UserResponse{
		ID:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusCreated, responseUser)
}
