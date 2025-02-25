package main

import (
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/controller"
	"github.com/OruamC/golang-crud-api/src/controller/routes"
	"github.com/OruamC/golang-crud-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	serviceDomain := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(serviceDomain)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting application", err)
	}
}
