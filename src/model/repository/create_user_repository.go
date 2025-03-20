package repository

import (
	"context"
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/configuration/rest_err"
	"github.com/OruamC/golang-crud-api/src/model"
	"github.com/OruamC/golang-crud-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error getting json value", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	logger.Info("User created successfully", zap.String("journey", "createUser"))

	value.ID = result.InsertedID.(bson.ObjectID)
	return converter.ConvertEntityToDomain(value), nil
}
