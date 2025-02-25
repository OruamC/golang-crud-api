package service

import (
	"github.com/OruamC/golang-crud-api/src/configuration/rest_err"
	"github.com/OruamC/golang-crud-api/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
