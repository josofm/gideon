package mock

import (
	"github.com/josofm/gideon/model"
)

type RepositoryMock struct {
	User  model.User
	Err   error
	Email string
}

func (r *RepositoryMock) Login(email, pass string) (model.User, error) {
	return r.User, r.Err
}

func (r *RepositoryMock) CreateUser(user model.User) (string, error) {
	return r.Email, r.Err
}

func (r *RepositoryMock) GetUser(id float64) (model.User, error) {
	return r.User, r.Err
}
