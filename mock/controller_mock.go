package mock

import (
	"github.com/MagicTheGathering/mtg-sdk-go"
	"github.com/josofm/gideon/model"
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

func (c *ControllerMock) GetUser(id float64) (model.User, error) {
	return c.User, c.ErrGetUser
}

func (c *ControllerMock) GetCardByName(name string) ([]*mtg.Card, error) {
	return c.Cards, c.ErrGetCards
}

func (c *ControllerMock) CreateDeck(deck model.Deck, userId float64) (string, error) {
	return c.DeckName, c.ErrGetDeck
}
