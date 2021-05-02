package mock

import (
	"github.com/josofm/gideon/model"
	"github.com/josofm/mtg-sdk-go"
)

type ControllerMock struct {
	Token       map[string]interface{}
	TokenModel  model.Token
	Email       string
	Err         error
	ErrGetUser  error
	User        model.User
	Cards       []*mtg.Card
	ErrGetCards error
	DeckName    string
	ErrGetDeck  error
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

func (c *ControllerMock) GetUser(id uint) (model.User, error) {
	return c.User, c.ErrGetUser
}

func (c *ControllerMock) GetCardByName(name string) ([]*mtg.Card, error) {
	return c.Cards, c.ErrGetCards
}

func (c *ControllerMock) CreateDeck(deck model.Deck, userId uint) (string, error) {
	return c.DeckName, c.ErrGetDeck
}

func (c *ControllerMock) DeleteUser(id uint) error {
	return c.Err
}

func (c *ControllerMock) UpdateUser(user model.User) error {
	return c.Err
}
