package repository

import (
	"context"
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/configuration/rest_err"
	"github.com/OruamC/golang-crud-api/src/model"
	"os"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomain.GetJsonValue()
	if err != nil {
		logger.Error("Error getting json value", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error getting json value", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}
