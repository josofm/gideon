package test

import (
	"github.com/josofm/gideon/model"
)

type ControllerMock struct {
	User model.User
	Err  error
}

func (c *ControllerMock) Login(name, pass string) (model.User, error) {
	return c.User, c.Err
}
