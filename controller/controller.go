package controller

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/josofm/gideon/helper"
	"github.com/josofm/gideon/model"
	"github.com/josofm/mtg-sdk-go"
)

type Controller struct {
	repository Repository
	clock      TimeClock
	secret     string
}

type Repository interface {
	Login(email, pass string) (model.User, error)
	CreateUser(user model.User) (string, error)
	GetUser(id uint) (model.User, error)
	CreateDeck(deck model.Deck) (string, error)
	UpdateUser(user model.User) error
	DeleteUser(user model.User) error
}

type TimeClock interface {
	Now() time.Time
	Add(t time.Time, d time.Duration) time.Time
}

func NewController(r Repository, c TimeClock) *Controller {
	s := os.Getenv("SECRET_JWT")
	return &Controller{
		repository: r,
		clock:      c,
		secret:     s,
	}
}

func (c *Controller) Login(email, pass string) (map[string]interface{}, error) {
	var user model.User
	if email == "" || pass == "" {
		return nil, errors.New("some error temporary")
	}
	user, err := c.repository.Login(email, pass)
	if err != nil {
		return nil, err
	}

	t, err := c.createToken(user)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (c *Controller) createToken(user model.User) (map[string]interface{}, error) {
	expiresAt := c.clock.Now().Add(time.Hour * 1).Unix()
	tk := &model.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(c.secret))
	if err != nil {
		return nil, err
	}
	t := map[string]interface{}{
		"token": tokenString,
	}
	return t, nil
}

func (c *Controller) CreateUser(user model.User) (string, error) {
	if !helper.UserHasAllFields(user) {
		log.Print("[createUser] missing parameters")
		return "", errors.New("Missing parameters")
	}
	email, err := c.repository.CreateUser(user)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (c *Controller) GetToken(header string) (model.Token, error) {
	tk := &model.Token{}

	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secret), nil
	})
	if err != nil {
		return model.Token{}, err
	}
	return *tk, err
}

func (c *Controller) GetUser(id uint) (model.User, error) {
	var user model.User
	user, err := c.repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *Controller) DeleteUser(id uint) error {
	var user model.User
	user, err := c.repository.GetUser(id)
	if err != nil {
		return err
	}
	return c.repository.DeleteUser(user)
}

func (c *Controller) UpdateUser(user model.User) error {
	if err := c.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetCardByName(name string) ([]*mtg.Card, error) {
	q := mtg.NewQuery()
	cards, err := q.Where(mtg.CardName, name).All()
	if err != nil {
		return cards, err
	}
	return cards, nil
}

func (c *Controller) CreateDeck(deck model.Deck, userId uint) (string, error) {
	deck.Owner.ID = userId
	if !helper.IsValidCommander(deck.Commander.Card.MultiverseId) {
		return "", errors.New("A commander must be a legendary card")
	}
	deckName, err := c.repository.CreateDeck(deck)
	if err != nil {
		return "", nil
	}
	return deckName, nil
}
