package demo_one

import (
	"strconv"
)

//go:generate mockgen -destination mock/user_service_mock.go -package mock -source user_service.go

type UserSrv interface {
	GetUserById(id int) string
	Update(id int, data interface{}) (string, error)
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserById(id int) string {
	return strconv.Itoa(id)
}

func (u *UserService) Update(id int, data interface{}) (string, error) {
	str := u.GetUserById(id)
	return str, nil
}
