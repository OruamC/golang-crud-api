package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/OruamC/golang-crud-api/src/configuration/logger"
	"github.com/OruamC/golang-crud-api/src/configuration/rest_err"
	"github.com/OruamC/golang-crud-api/src/model"
	"github.com/OruamC/golang-crud-api/src/model/repository/entity"
	"github.com/OruamC/golang-crud-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Info(errorMessage,
				zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find user by email",
			err,
			zap.String("journey", "findUserByEmail"))
		errorMessage := "Error trying to find user by email"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByID(
	ID string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: ID}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", ID)
			logger.Info(errorMessage,
				zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find user by ID",
			err,
			zap.String("journey", "findUserByID"))
		errorMessage := "Error trying to find user by ID"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(userEntity), nil
}
