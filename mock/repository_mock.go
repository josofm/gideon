package mock

import (
	"github.com/josofm/gideon/model"
)

type RespositoryMock struct {
	User model.User
	Err  error
}

func (r *RespositoryMock) Login(email, pass string) (model.User, error) {
	return r.User, r.Err
}
