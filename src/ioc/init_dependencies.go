package ioc

import (
	"github.com/OruamC/golang-crud-api/src/controller"
	"github.com/OruamC/golang-crud-api/src/model/repository"
	"github.com/OruamC/golang-crud-api/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	serviceDomain := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(serviceDomain)
}
