package mock

import (
	"github.com/josofm/gideon/model"
)

type ControllerMock struct {
	Token      map[string]interface{}
	TokenModel model.Token
	Email      string
	Err        error
	User       model.User
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

func (c *ControllerMock) GetUser(id float64) (model.User, error) {
	return c.User, c.Err
}
