package main

import (
	"context"
	"github.com/OruamC/golang-crud-api/src/configuration/database/mongodb"
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/controller"
	"github.com/OruamC/golang-crud-api/src/controller/routes"
	"github.com/OruamC/golang-crud-api/src/model/repository"
	"github.com/OruamC/golang-crud-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}
	defer database.Client().Disconnect(context.Background())

	// Init dependencies
	repo := repository.NewUserRepository(database)
	serviceDomain := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(serviceDomain)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error starting application", err)
	}
}
