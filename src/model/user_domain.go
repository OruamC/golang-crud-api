package model

type userDomain struct {
	iD       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetId() string {
	return ud.iD
}

func (ud *userDomain) SetId(id string) {
	ud.iD = id
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
