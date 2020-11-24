package mock

import (
	"github.com/josofm/gideon/model"
)

type ControllerMock struct {
	Token      map[string]interface{}
	TokenModel model.Token
	Email      string
	Err        error
}

func (c *ControllerMock) Login(name, pass string) (map[string]interface{}, error) {
	return c.Token, c.Err
}

func (c *ControllerMock) CreateUser(user model.User) (string, error) {
	return c.Email, c.Err
}

func (c *ControllerMock) GetToken(header string) (model.Token, error) {
	return c.TokenModel, c.Err
}
