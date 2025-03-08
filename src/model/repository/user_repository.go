package repository

import (
	"github.com/OruamC/golang-crud-api/src/configuration/rest_err"
	"github.com/OruamC/golang-crud-api/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
