package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetId(string)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) GetId() string {
	return ud.iD
}

func (ud *userDomain) SetId(id string) {
	ud.iD = id
}

type userDomain struct {
	iD       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
