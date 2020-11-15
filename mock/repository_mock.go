package mock

import (
	"github.com/josofm/gideon/model"
)

type RespositoryMock struct {
	User  model.User
	Err   error
	Email string
}

func (r *RespositoryMock) Login(email, pass string) (model.User, error) {
	return r.User, r.Err
}

func (r *RespositoryMock) CreateUser(user model.User) (string, error) {
	return r.Email, r.Err
}
